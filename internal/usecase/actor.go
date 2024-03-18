package usecase

import (
	"context"
	"filmoteca/internal/entity"
	"filmoteca/internal/repository"
)

type ActorUsecase struct {
	repos repository.Repositories
}

func NewActorUsecase(repos repository.Repositories) ActorUsecase {
	return ActorUsecase{
		repos: repos,
	}
}

func (a ActorUsecase) Create(ctx context.Context, actor entity.Actor) error {
	return a.repos.Actors.Create(ctx, actor)
}

func (a ActorUsecase) ReadOne(ctx context.Context, id string) (entity.Actor, error) {
	return a.repos.Actors.ReadOne(ctx, id)
}

func (a ActorUsecase) ReadAll(ctx context.Context) ([]entity.Actor, error) {
	return a.repos.Actors.ReadAll(ctx)
}

func (a ActorUsecase) UpdateInfo(ctx context.Context, id string, actor entity.Actor) (entity.Actor, error) {
	return a.repos.Actors.Update(ctx, id, actor)
}

func (a ActorUsecase) Delete(ctx context.Context, id string) (entity.Actor, error) {
	return a.repos.Actors.DeleteOne(ctx, id)
}

func (a ActorUsecase) DeleteInfo(ctx context.Context, id string, fields []string) (entity.Actor, error) {
	return a.repos.Actors.DeleteInfo(ctx, id, fields)
}
