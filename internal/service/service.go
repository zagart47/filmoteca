package service

import (
	"context"
	"filmoteca/internal/entity"
	"filmoteca/internal/repository"
	"filmoteca/internal/usecase"
)

type Actors interface {
	Create(ctx context.Context, actor entity.Actor) error
	ReadOne(ctx context.Context, id string) (entity.Actor, error)
	ReadAll(ctx context.Context) ([]entity.Actor, error)
	Update(ctx context.Context, id string, actor entity.Actor) (entity.Actor, error)
	Delete(ctx context.Context, id string, options []string) (entity.Actor, error)
}

type Movies interface {
	Create(ctx context.Context, Movie entity.Movie, id string) error
	ReadOne(ctx context.Context, id string) (entity.Movie, error)
	ReadAll(ctx context.Context, options entity.Options) ([]entity.Movie, error)
	Update(ctx context.Context, id string, Movie entity.Movie) (entity.Movie, error)
	Delete(ctx context.Context, id string, options []string) (entity.Movie, error)
}

type Users interface {
	Get(ctx context.Context) ([]entity.User, error)
}

type Services struct {
	Actors Actors
	Movies Movies
	Users  Users
}

func NewServices(repos repository.Repositories) Services {
	usecases := usecase.NewUsecases(repos)
	userService := NewUserService(usecases)
	actorService := NewActorService(usecases)
	movieService := NewMovieService(usecases)
	return Services{
		Actors: actorService,
		Movies: movieService,
		Users:  userService,
	}
}
