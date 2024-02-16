package service

import (
	"context"
	"my-message-app/internal/domain"
)

type MessageService interface {
	CreateMessage(ctx context.Context, message domain.Message) (*domain.Message, error)
	GetMessages(ctx context.Context) ([]domain.Message, error)
}

type Services struct {
	MessageService MessageService
}

func NewServices(messageService MessageService) *Services {
	return &Services{
		MessageService: messageService,
	}
}
