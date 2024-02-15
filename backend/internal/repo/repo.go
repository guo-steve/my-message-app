package repo

import (
	"context"

	"my-message-app/internal/domain"
)

type Repository interface {
	PostMessage(ctx context.Context, message domain.Message) (*domain.Message, error)
	GetMessages(ctx context.Context) ([]domain.Message, error)
}
