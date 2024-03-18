package postgresql

import (
	"context"
	"errors"
	"filmoteca/internal/entity"
	"filmoteca/pkg/logger"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
)

type MovieRepo struct {
	Db Client
}

func NewMovieRepo(db Client) *MovieRepo {
	return &MovieRepo{
		Db: db,
	}
}

func (m MovieRepo) Create(ctx context.Context, movie entity.Movie, actorId string) error {
	query := `INSERT INTO public.movies (title, description, release_date, rating, is_del)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING (id)`
	if err := m.Db.QueryRow(ctx, query, movie.Title, movie.Description, movie.ReleaseDate, movie.Rating, false).Scan(&movie.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			logger.Log.Error("sql error:", pgErr.Message, "detail:", pgErr.Detail, "where:", pgErr.Where, "")
			return pgErr
		} else if err != nil {
			return err
		}
	}
	q := `INSERT INTO public.movies_actors (movie_id, actor_id, is_del)
		VALUES ($1, $2, $3)
		RETURNING (id)`
	if err := m.Db.QueryRow(ctx, q, movie.ID, actorId, false).Scan(&movie.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			logger.Log.Warn("sql error:", pgErr.Message, "detail:", pgErr.Detail, "where:", pgErr.Where, "")
			return pgErr
		} else if err != nil {
			return err
		}
	}
	logger.Log.Info("movie added to db", "id", movie.ID)
	return nil
}

func (m MovieRepo) ReadOne(ctx context.Context, id string) (entity.Movie, error) {
	movie := entity.NewMovie()
	pq := `	SELECT id, title, description, release_date, rating
	FROM public.movies
	WHERE id = $1 AND is_del = false `
	rows, err := m.Db.Query(ctx, pq, id)
	if err != nil {
		return movie, err
	}
	for rows.Next() {
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				logger.Log.Error("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			}
			return entity.Movie{}, err
		}
	}
	logger.Log.Info("movie getting from db successful", "id", id)

	return movie, nil
}

func (m MovieRepo) ReadAll(ctx context.Context, Options entity.Options) ([]entity.Movie, error) {
	q := sq.Select("id", "title", "description", "release_date", "rating").Where("is_del = false").From("public.movies")
	if Options.Title != "" {
		q = q.Where(fmt.Sprint("title LIKE '%", Options.Title, "%' AND is_del = false"))
	}
	if Options.Actor != "" {
		q = sq.Select("m.id", "m.title", "m.description", "m.release_date", "m.rating").
			From("public.movies m").
			Join("public.movies_actors ma ON m.id = ma.movie_id").
			Join("public.actors a ON ma.actor_id = a.id").
			Where(fmt.Sprint("a.name LIKE '%", Options.Actor, "%' AND m.is_del = false"))
	}
	if Options.Field != "" {
		if Options.Order != "" {
			q = q.OrderBy(fmt.Sprintf("%s %s", Options.Field, Options.Order))
		} else {
			q = q.OrderBy(fmt.Sprintf("%s DESC", Options.Field))
		}
	} else {
		q = q.OrderBy(fmt.Sprintf("rating DESC"))
	}
	sql, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := m.Db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	var movies = make([]entity.Movie, 0)
	for rows.Next() {
		var movie entity.Movie
		if err = rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				logger.Log.Error("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			}
			return nil, err
		}
		movies = append(movies, movie)
	}
	logger.Log.Info("movies getting from db successful")
	return movies, nil
}

func (m MovieRepo) Update(ctx context.Context, id string, movie entity.Movie) (entity.Movie, error) {
	q := sq.Update("movies").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")
	if movie.Title != "" {
		q = q.Set("title", movie.Title)
	}
	if movie.Description != "" {
		q = q.Set("description", movie.Description)
	}
	if movie.ReleaseDate != "" {
		q = q.Set("release_date", movie.ReleaseDate)
	}
	if movie.Rating != 0 {
		q = q.Set("rating", movie.Rating)
	}
	sql, args, err := q.ToSql()
	if err != nil {
		logger.Log.Error("updating data error", err.Error(), "")
		return entity.Movie{}, err
	}
	err = m.Db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		logger.Log.Error("updating data error", err.Error(), "")
		return entity.Movie{}, err
	}
	logger.Log.Info("movie editing successful", "id", id)
	return m.ReadOne(ctx, id)
}

func (m MovieRepo) DeleteOne(ctx context.Context, id string) (entity.Movie, error) {
	q, args, err := sq.Update("public.movies").
		Set("is_del", true).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).ToSql()
	_, err = m.Db.Exec(ctx, q, args...)
	if err != nil {
		logger.Log.Error("movie deleting error", err.Error(), "")
		return entity.Movie{}, err
	}
	q, args, err = sq.Update("public.movies_actors").
		Set("is_del", true).
		Where(sq.Eq{"movie_id": id}).
		PlaceholderFormat(sq.Dollar).ToSql()
	_, err = m.Db.Exec(ctx, q, args...)
	if err != nil {
		logger.Log.Error("movie deleting error", err.Error(), "")
		return entity.Movie{}, err
	}
	logger.Log.Info("movie deleted", "id", id)

	return entity.Movie{}, nil
}

func (m MovieRepo) DeleteInfo(ctx context.Context, id string, fields []string) (entity.Movie, error) {
	q := sq.Update("actors").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")
	for _, v := range fields {
		q = q.Set(v, nil)
	}
	sql, args, err := q.ToSql()
	if err != nil {
		logger.Log.Error("updating data error", err.Error(), "")
		return entity.Movie{}, err
	}
	err = m.Db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		logger.Log.Error("updating data error", err.Error(), "")
		return entity.Movie{}, err
	}
	logger.Log.Info("movie info editing success", "id", id)
	return m.ReadOne(ctx, id)
}
