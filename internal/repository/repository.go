package repository

import (
	"context"
	"filmoteca/internal/entity"
	"filmoteca/internal/repository/postgresql"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Actors interface {
	Create(context.Context, entity.Actor) error
	ReadOne(context.Context, string) (entity.Actor, error)
	ReadAll(context.Context) ([]entity.Actor, error)
	Update(context.Context, string, entity.Actor) (entity.Actor, error)
	DeleteOne(context.Context, string) (entity.Actor, error)
	DeleteInfo(context.Context, string, []string) (entity.Actor, error)
}

type Users interface {
	Get(ctx context.Context) ([]entity.User, error)
}

type Movies interface {
	Create(context.Context, entity.Movie, string) error
	ReadOne(context.Context, string) (entity.Movie, error)
	ReadAll(context.Context, entity.Options) ([]entity.Movie, error)
	Update(context.Context, string, entity.Movie) (entity.Movie, error)
	DeleteOne(context.Context, string) (entity.Movie, error)
	DeleteInfo(context.Context, string, []string) (entity.Movie, error)
}

type Repositories struct {
	Actors Actors
	Movies Movies
	Users  Users
}

func NewRepositories(db *pgxpool.Pool) Repositories {
	return Repositories{
		Actors: postgresql.NewActorRepo(db),
		Movies: postgresql.NewMovieRepo(db),
		Users:  postgresql.NewUserRepo(db),
	}
}
