package usecase

import (
	"context"
	"filmoteca/internal/entity"
	"filmoteca/internal/repository"
)

type MovieUsecase struct {
	Repos repository.Repositories
}

func NewMovieUsecase(repos repository.Repositories) MovieUsecase {
	return MovieUsecase{
		Repos: repos,
	}
}

func (m MovieUsecase) Create(ctx context.Context, Movie entity.Movie, actorId string) error {
	return m.Repos.Movies.Create(ctx, Movie, actorId)
}

func (m MovieUsecase) ReadOne(ctx context.Context, id string) (entity.Movie, error) {
	return m.Repos.Movies.ReadOne(ctx, id)
}

func (m MovieUsecase) ReadAll(ctx context.Context, Options entity.Options) ([]entity.Movie, error) {
	return m.Repos.Movies.ReadAll(ctx, Options)
}

func (m MovieUsecase) UpdateInfo(ctx context.Context, id string, Movie entity.Movie) (entity.Movie, error) {
	return m.Repos.Movies.Update(ctx, id, Movie)
}

func (m MovieUsecase) Delete(ctx context.Context, id string) (entity.Movie, error) {
	return m.Repos.Movies.DeleteOne(ctx, id)
}

func (m MovieUsecase) DeleteInfo(ctx context.Context, id string, fields []string) (entity.Movie, error) {
	return m.Repos.Movies.DeleteInfo(ctx, id, fields)
}
