package service

import (
	"context"

	"my-message-app/internal/domain"
	"my-message-app/internal/repo"
)

var _ MessageService = (*messageService)(nil)

type messageService struct {
	repo repo.MessageRepo
}

func NewMessageService(repo repo.MessageRepo) *messageService {
	return &messageService{
		repo: repo,
	}
}

func (s *messageService) CreateMessage(ctx context.Context, message domain.Message) (*domain.Message, error) {
	return s.repo.CreateMessage(ctx, message)
}

func (s *messageService) GetMessages(ctx context.Context, createdBy string) ([]domain.Message, error) {
	return s.repo.GetMessages(ctx, createdBy)
}

func (s *messageService) UpdateMessage(ctx context.Context, message domain.Message) (*domain.Message, error) {
	return s.repo.UpdateMessage(ctx, message)
}

func (s *messageService) DeleteMessage(ctx context.Context, messageID, userID string) error {
	return s.repo.DeleteMessage(ctx, messageID, userID)
}
