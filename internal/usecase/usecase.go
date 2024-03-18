package usecase

import (
	"context"
	"filmoteca/internal/entity"
	"filmoteca/internal/repository"
)

type Actors interface {
	Create(ctx context.Context, actor entity.Actor) error
	ReadOne(ctx context.Context, id string) (entity.Actor, error)
	ReadAll(ctx context.Context) ([]entity.Actor, error)
	UpdateInfo(ctx context.Context, id string, actor entity.Actor) (entity.Actor, error)
	Delete(ctx context.Context, id string) (entity.Actor, error)
	DeleteInfo(ctx context.Context, id string, fields []string) (entity.Actor, error)
}

type Movies interface {
	Create(ctx context.Context, movie entity.Movie, actorId string) error
	ReadOne(ctx context.Context, id string) (entity.Movie, error)
	ReadAll(ctx context.Context, options entity.Options) ([]entity.Movie, error)
	UpdateInfo(ctx context.Context, id string, movie entity.Movie) (entity.Movie, error)
	Delete(ctx context.Context, id string) (entity.Movie, error)
	DeleteInfo(ctx context.Context, id string, fields []string) (entity.Movie, error)
}

type Users interface {
	Get(ctx context.Context) ([]entity.User, error)
}

type Usecases struct {
	Actors Actors
	Movies Movies
	Users  Users
}

func NewUsecases(repos repository.Repositories) Usecases {
	actorUsecase := NewActorUsecase(repos)
	movieUsecase := NewMovieUsecase(repos)
	userUsecase := NewUserUsecase(repos)
	return Usecases{
		Actors: actorUsecase,
		Movies: movieUsecase,
		Users:  userUsecase,
	}
}
