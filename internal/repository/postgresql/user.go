package postgresql

import (
	"context"
	"errors"
	"filmoteca/internal/entity"
	"filmoteca/pkg/logger"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepo struct {
	db Client
}

func NewUserRepo(db Client) *UserRepo {
	return &UserRepo{db: db}
}

func (r UserRepo) Get(ctx context.Context) ([]entity.User, error) {
	user := entity.NewUser()
	var users []entity.User
	pq := `	SELECT name FROM public.users WHERE role = $1`
	rows, err := r.db.Query(ctx, pq, "admin")
	if err != nil {
		return users, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Name); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				logger.Log.Error("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			}
			return []entity.User{}, err
		}
		users = append(users, user)
	}
	return users, nil
}
