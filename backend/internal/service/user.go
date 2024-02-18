package service

import (
	"context"

	"my-message-app/internal/domain"
	"my-message-app/internal/repo"
)

type userService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetUserByID(
	ctx context.Context,
	id string,
) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}
