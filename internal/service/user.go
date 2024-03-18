package service

import (
	"context"
	"filmoteca/internal/entity"
	"filmoteca/internal/usecase"
)

type UserService struct {
	usecase usecase.Usecases
	Users   Users
}

func NewUserService(u usecase.Usecases) UserService {
	return UserService{
		usecase: u,
	}
}

func (a UserService) Get(ctx context.Context) ([]entity.User, error) {
	return a.usecase.Users.Get(ctx)
}
