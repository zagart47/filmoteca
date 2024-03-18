package postgresql

import (
	"context"
	"errors"
	"filmoteca/internal/entity"
	"filmoteca/pkg/logger"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
)

type ActorRepo struct {
	Db Client
}

func NewActorRepo(db Client) *ActorRepo {
	return &ActorRepo{Db: db}
}

func (r ActorRepo) Create(ctx context.Context, actor entity.Actor) error {
	query := `INSERT INTO public.actors (name, gender, birthdate, is_del)
		VALUES ($1, $2, $3, $4)
		RETURNING (id)`
	if err := r.Db.QueryRow(ctx, query, actor.Name, actor.Birthdate, actor.Gender, false).Scan(&actor.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			logger.Log.Error("sql error:", pgErr.Message, "detail:", pgErr.Detail, "where:", pgErr.Where)
			return pgErr
		} else if err != nil {
			return err
		}
	}
	logger.Log.Info("actor added to db", "id", actor.ID)
	return nil
}

func (r ActorRepo) ReadOne(ctx context.Context, id string) (entity.Actor, error) {
	actor := entity.NewActor()
	pq := `	SELECT a.id, a.name, a.gender, a.birthdate, a.is_del, m.title
			FROM (SELECT id, name, gender, birthdate, is_del
			FROM public.actors
			WHERE id = $1 AND is_del = false) a
			LEFT JOIN (SELECT movies.title, movies.is_del, movies_actors.actor_id
			FROM movies
			JOIN movies_actors ON movies.id = movies_actors.movie_id
			WHERE movies_actors.actor_id = $1 AND movies_actors.is_del = false) m ON a.id = m.actor_id
			WHERE a.is_del = false AND m.is_del = false`
	rows, err := r.Db.Query(ctx, pq, id)
	if err != nil {
		return actor, err
	}
	var movie string
	for rows.Next() {
		var isDel bool
		if err := rows.Scan(&actor.ID, &actor.Name, &actor.Gender, &actor.Birthdate, &isDel, &movie); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				logger.Log.Error("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			}
			return entity.Actor{}, err
		}
		actor.Movies = append(actor.Movies, movie)
	}
	logger.Log.Info("actor getting from db successful", "id", id)
	return actor, nil
}

func (r ActorRepo) ReadAll(ctx context.Context) ([]entity.Actor, error) {
	pq := `
		SELECT id, name, gender, birthdate FROM public.actors
		WHERE is_del = false
		ORDER BY id
		`
	rows, err := r.Db.Query(ctx, pq)
	if err != nil {
		return nil, err
	}
	var actors = make([]entity.Actor, 0)
	for rows.Next() {
		var actor entity.Actor
		if err = rows.Scan(&actor.ID, &actor.Name, &actor.Gender, &actor.Birthdate); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				logger.Log.Error("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			}
			return nil, err
		}
		cq := `	SELECT movies.title
				FROM movies
				JOIN movies_actors ON movies.id = movies_actors.movie_id
				JOIN actors ON movies_actors.actor_id = actors.id
				WHERE actors.id = $1 AND movies.is_del = false
				ORDER BY movies.release_date
		`
		actorMovies, err := r.Db.Query(ctx, cq, actor.ID)
		if err != nil {
			return nil, err
		}
		var m string
		for actorMovies.Next() {
			if err = actorMovies.Scan(&m); err != nil {
				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) {
					logger.Log.Error("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
				}
				return nil, err
			}
			actor.Movies = append(actor.Movies, m)
		}
		actors = append(actors, actor)
	}
	logger.Log.Info("actors getting from db successful")
	return actors, nil
}

func (r ActorRepo) Update(ctx context.Context, id string, actor entity.Actor) (entity.Actor, error) {
	q := sq.Update("actors").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")
	if actor.Name != "" {
		q = q.Set("name", actor.Name)
	}
	if actor.Gender != "" {
		q = q.Set("gender", actor.Gender)
	}
	if actor.Birthdate != "" {
		q = q.Set("birthdate", actor.Birthdate)
	}
	sql, args, err := q.ToSql()
	if err != nil {
		logger.Log.Error("updating data error", err.Error(), "")
		return entity.Actor{}, err
	}
	err = r.Db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		logger.Log.Error("updating data error", err.Error(), "")
		return entity.Actor{}, err
	}
	logger.Log.Info("actor editing successful", "id", id)
	return r.ReadOne(ctx, id)
}

func (r ActorRepo) DeleteOne(ctx context.Context, id string) (entity.Actor, error) {
	q, args, err := sq.Update("public.actors").
		Set("is_del", true).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).ToSql()
	_, err = r.Db.Exec(ctx, q, args...)
	if err != nil {
		logger.Log.Error("actor deleting error", err.Error(), "")
		return entity.Actor{}, err
	}
	q, args, err = sq.Update("public.movies_actors").
		Set("is_del", true).
		Where(sq.Eq{"actor_id": id}).
		PlaceholderFormat(sq.Dollar).ToSql()
	_, err = r.Db.Exec(ctx, q, args...)
	if err != nil {
		logger.Log.Error("actor deleting error", err.Error(), "")
		return entity.Actor{}, err
	}
	logger.Log.Info("actor deleted", "id", id)
	return entity.Actor{}, nil
}
func (r ActorRepo) DeleteInfo(ctx context.Context, id string, fields []string) (entity.Actor, error) {
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
		return entity.Actor{}, err
	}
	err = r.Db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		logger.Log.Error("updating data error", err.Error(), "")
		return entity.Actor{}, err
	}
	logger.Log.Info("actor info editing success", "id", id)
	return r.ReadOne(ctx, id)
}
