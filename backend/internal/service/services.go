package service

import (
	"context"

	"my-message-app/internal/domain"
)

type AuthService interface {
	Register(ctx context.Context, user domain.User) (*domain.User, error)
	Login(ctx context.Context, creds domain.Credentials) (string, error)
	Authenticate(ctx context.Context, tokenString string) (*domain.User, error)
	Logout(ctx context.Context, tokenString string) error
}

type MessageService interface {
	CreateMessage(ctx context.Context, message domain.Message) (*domain.Message, error)
	GetMessages(ctx context.Context) ([]domain.Message, error)
}

type Services struct {
	MessageService MessageService
	AuthService    AuthService
}

func NewServices(messageService MessageService, authService AuthService) *Services {
	return &Services{
		MessageService: messageService,
		AuthService:    authService,
	}
}
