package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"my-message-app/internal/domain"
	"my-message-app/internal/repo"
)

func TestNewMessageService(t *testing.T) {
	type args struct {
		repo repo.MessageRepo
	}
	tests := []struct {
		name string
		args args
		want *messageService
	}{
		{
			name: "TestNewMessageService",
			args: args{
				repo: &mockMessageRepo{},
			},
			want: &messageService{
				repo: &mockMessageRepo{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessageService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessageService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messageService_CreateMessage(t *testing.T) {
	type fields struct {
		repo repo.MessageRepo
	}
	type args struct {
		ctx     context.Context
		message domain.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Message
		wantErr bool
	}{
		{
			name: "TestCreateMessage",
			fields: fields{
				repo: &mockMessageRepo{},
			},
			args: args{
				ctx: context.Background(),
				message: domain.Message{
					Content: "Hello",
				},
			},
			want: &domain.Message{
				ID:      "1",
				Content: "Hello",
			},
			wantErr: false,
		},
		{
			name: "TestCreateMessageError",
			fields: fields{
				repo: &mockMessageRepo{
					hasError: true,
				},
			},
			args: args{
				ctx: context.Background(),
				message: domain.Message{
					Content: "Hello",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &messageService{
				repo: tt.fields.repo,
			}
			got, err := s.CreateMessage(tt.args.ctx, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("messageService.CreateMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messageService.CreateMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messageService_GetMessages(t *testing.T) {
	type fields struct {
		repo repo.MessageRepo
	}
	type args struct {
		ctx       context.Context
		createdBy string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Message
		wantErr bool
	}{
		{
			name: "TestGetMessages",
			fields: fields{
				repo: &mockMessageRepo{},
			},
			args: args{
				ctx: context.Background(),
			},
			want: []domain.Message{
				{
					ID:      "1",
					Content: "Hello",
				},
			},
			wantErr: false,
		},
		{
			name: "TestGetMessagesError",
			fields: fields{
				repo: &mockMessageRepo{
					hasError: true,
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &messageService{
				repo: tt.fields.repo,
			}
			got, err := s.GetMessages(tt.args.ctx, tt.args.createdBy)
			if (err != nil) != tt.wantErr {
				t.Errorf("messageService.GetMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messageService.GetMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockMessageRepo struct {
	hasError bool
}

var _ repo.MessageRepo = &mockMessageRepo{}

func (m *mockMessageRepo) CreateMessage(ctx context.Context, message domain.Message) (*domain.Message, error) {
	if m.hasError {
		return nil, errors.New("error")
	}
	return &domain.Message{
		ID:      "1",
		Content: message.Content,
	}, nil
}

func (m *mockMessageRepo) GetMessages(ctx context.Context, createdBy string) ([]domain.Message, error) {
	if m.hasError {
		return nil, errors.New("error")
	}

	return []domain.Message{
		{
			ID:      "1",
			Content: "Hello",
		},
	}, nil
}

func (m *mockMessageRepo) UpdateMessage(ctx context.Context, message domain.Message) (*domain.Message, error) {
	if m.hasError {
		return nil, errors.New("error")
	}
	return &message, nil
}

func (m *mockMessageRepo) DeleteMessage(ctx context.Context, mID, userID string) error {
	if m.hasError {
		return errors.New("error")
	}
	return nil
}
