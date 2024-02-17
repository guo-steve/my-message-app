package repo

import (
	"context"

	"my-message-app/internal/domain"
)

type MessageRepo interface {
	CreateMessage(ctx context.Context, message domain.Message) (*domain.Message, error)
	GetMessages(ctx context.Context) ([]domain.Message, error)
}

type UserRepo interface {
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
}

type InvalidTokenRepo interface {
	CreateInvalidToken(ctx context.Context, tokenString string) error
	FindInvalidToken(ctx context.Context, tokenString string) (*domain.InvalidToken, error)
}
