package service

import (
	"context"
	"filmoteca/internal/entity"
	"filmoteca/internal/usecase"
)

type ActorService struct {
	usecase usecase.Usecases
	Actors  Actors
}

func NewActorService(u usecase.Usecases) ActorService {
	return ActorService{
		usecase: u,
	}
}

func (a ActorService) Create(ctx context.Context, actor entity.Actor) error {
	return a.usecase.Actors.Create(ctx, actor)
}

func (a ActorService) ReadOne(ctx context.Context, id string) (entity.Actor, error) {
	return a.usecase.Actors.ReadOne(ctx, id)
}

func (a ActorService) ReadAll(ctx context.Context) ([]entity.Actor, error) {
	return a.usecase.Actors.ReadAll(ctx)
}

func (a ActorService) Update(ctx context.Context, id string, actor entity.Actor) (entity.Actor, error) {
	return a.usecase.Actors.UpdateInfo(ctx, id, actor)
}

func (a ActorService) Delete(ctx context.Context, id string, options []string) (entity.Actor, error) {
	if len(options) == 0 {
		return a.usecase.Actors.Delete(ctx, id)
	}
	return a.usecase.Actors.DeleteInfo(ctx, id, options)
}
