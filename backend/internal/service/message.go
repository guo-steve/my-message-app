package service

import (
	"context"

	"my-message-app/internal/domain"
	"my-message-app/internal/repo"
)

var _ MessageService = (*messageService)(nil)

type messageService struct {
	repo repo.Repository
}

func NewMessageService(repo repo.Repository) *messageService {
	return &messageService{
		repo: repo,
	}
}

func (s *messageService) CreateMessage(ctx context.Context, message domain.Message) (*domain.Message, error) {
	return s.repo.CreateMessage(ctx, message)
}

func (s *messageService) GetMessages(ctx context.Context) ([]domain.Message, error) {
	return s.repo.GetMessages(ctx)
}
