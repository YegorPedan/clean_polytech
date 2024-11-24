package user

import (
	"clean-polytech/internal/domain/model"
	"clean-polytech/internal/domain/repository"
	"context"
)

type GetUsersUse struct {
	repo repository.UserRepository
}

func NewGetUsersUseCase(repo repository.UserRepository) *GetUsersUse {
	return &GetUsersUse{
		repo: repo,
	}
}

func (g *GetUsersUse) Execute(context.Context) ([]*model.User, error) {
	return g.repo.GetAllUsers()
}
