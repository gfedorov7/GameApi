package service

import (
	"GameApi/internal/model"
	"GameApi/internal/repository"
	"context"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) FindById(ctx context.Context, id int64) (*model.User, error) {
	return service.repo.GetById(ctx, id)
}

func (service *UserService) FindAll(ctx context.Context) ([]model.User, error) {
	return service.repo.GetByAll(ctx)
}
