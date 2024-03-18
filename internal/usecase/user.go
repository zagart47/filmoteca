package usecase

import (
	"context"
	"filmoteca/internal/entity"
	"filmoteca/internal/repository"
)

type UserUsecase struct {
	repos repository.Repositories
}

func NewUserUsecase(repos repository.Repositories) UserUsecase {
	return UserUsecase{
		repos: repos,
	}
}

func (a UserUsecase) Get(ctx context.Context) ([]entity.User, error) {
	return a.repos.Users.Get(ctx)
}
