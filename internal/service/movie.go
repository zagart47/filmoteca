package service

import (
	"context"
	"filmoteca/internal/entity"
	"filmoteca/internal/usecase"
)

type MovieService struct {
	Movies  Movies
	usecase usecase.Usecases
}

func NewMovieService(usecase usecase.Usecases) MovieService {
	return MovieService{
		usecase: usecase,
	}
}

func (a MovieService) Create(ctx context.Context, movie entity.Movie, id string) error {
	return a.usecase.Movies.Create(ctx, movie, id)
}

func (a MovieService) ReadOne(ctx context.Context, id string) (entity.Movie, error) {
	return a.usecase.Movies.ReadOne(ctx, id)
}

func (a MovieService) ReadAll(ctx context.Context, Options entity.Options) ([]entity.Movie, error) {
	return a.usecase.Movies.ReadAll(ctx, Options)
}

func (a MovieService) Update(ctx context.Context, id string, movie entity.Movie) (entity.Movie, error) {
	return a.usecase.Movies.UpdateInfo(ctx, id, movie)
}

func (a MovieService) Delete(ctx context.Context, id string, Options []string) (entity.Movie, error) {
	if len(Options) == 0 {
		return a.usecase.Movies.Delete(ctx, id)
	}
	return a.usecase.Movies.DeleteInfo(ctx, id, Options)
}
