package api

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"my-message-app/internal/domain"
	"my-message-app/internal/service"
)

func TestNewHandler(t *testing.T) {
	services := service.NewServices(&mockMessageService{})

	type args struct {
		services *service.Services
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "TestNewHandler",
			args: args{
				services: services,
			},
			want: &Handler{
				services: services,
				logger:   slog.Default(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.services, slog.Default()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_PostMessage(t *testing.T) {
	req := func() *http.Request {
		return httptest.NewRequest(
			http.MethodPost,
			"/v1/messages",
			strings.NewReader(`{"content":"Hello"}`),
		)
	}

	type fields struct {
		services *service.Services
	}
	type args struct {
		resWtr http.ResponseWriter
		req    *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Message
		wantErr bool
	}{
		{
			name: "TestPostMessage",
			fields: fields{
				services: service.NewServices(&mockMessageService{}),
			},
			args: args{
				resWtr: httptest.NewRecorder(),
				req:    req(),
			},
			want: &domain.Message{
				ID:      "1",
				Content: "Hello",
			},
		},
		{
			name: "TestPostMessageError",
			fields: fields{
				services: service.NewServices(&mockMessageService{hasError: true}),
			},
			args: args{
				resWtr: httptest.NewRecorder(),
				req:    req(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				services: tt.fields.services,
				logger:   slog.Default(),
			}
			h.PostMessage(tt.args.resWtr, tt.args.req)

			res := tt.args.resWtr.(*httptest.ResponseRecorder).Result()
			resB, _ := io.ReadAll(res.Body)

			if !tt.wantErr && res.StatusCode != http.StatusCreated {
				t.Errorf("PostMessage() = %v, want %v", res.StatusCode, http.StatusCreated)
			}

			if tt.wantErr && res.StatusCode == http.StatusCreated {
				t.Errorf("WantErr %v PostMessage() = %v, want %v", tt.wantErr, res.StatusCode, http.StatusInternalServerError)
			}

			if !tt.wantErr {
				var got domain.Message
				_ = json.Unmarshal(resB, &got)
				if !reflect.DeepEqual(got, *tt.want) {
					t.Errorf("PostMessage() = %v, want %v", got, *tt.want)
				}
			}
			res.Body.Close()
		})
	}
}

func TestHandler_GetMessages(t *testing.T) {
	type fields struct {
		services *service.Services
	}
	type args struct {
		resWtr http.ResponseWriter
		req    *http.Request
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
				services: service.NewServices(&mockMessageService{}),
			},
			args: args{
				resWtr: httptest.NewRecorder(),
				req:    httptest.NewRequest(http.MethodGet, "/v1/messages", nil),
			},
			want: []domain.Message{
				{
					ID:      "1",
					Content: "Hello",
				},
				{
					ID:      "2",
					Content: "World",
				},
			},
		},
		{
			name: "TestGetMessagesError",
			fields: fields{
				services: service.NewServices(&mockMessageService{hasError: true}),
			},
			args: args{
				resWtr: httptest.NewRecorder(),
				req:    httptest.NewRequest(http.MethodGet, "/v1/messages", nil),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				services: tt.fields.services,
				logger:   slog.Default(),
			}
			h.GetMessages(tt.args.resWtr, tt.args.req)

			res := tt.args.resWtr.(*httptest.ResponseRecorder).Result()
			resB, _ := io.ReadAll(res.Body)

			if !tt.wantErr && res.StatusCode != http.StatusOK {
				t.Errorf("GetMessages() = %v, want %v", res.StatusCode, http.StatusOK)
			}

			if tt.wantErr && res.StatusCode == http.StatusOK {
				t.Errorf("WantErr %v GetMessages() = %v, want %v", tt.wantErr, res.StatusCode, http.StatusInternalServerError)
			}

			if !tt.wantErr {
				var got []domain.Message
				_ = json.Unmarshal(resB, &got)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetMessages() = %v, want %v", got, tt.want)
				}
			}
			res.Body.Close()
		})
	}
}

type mockMessageService struct {
	hasError bool
}

func (m *mockMessageService) CreateMessage(ctx context.Context, message domain.Message) (*domain.Message, error) {
	if m.hasError {
		return nil, errors.New("error")
	}
	return &domain.Message{
		ID:      "1",
		Content: message.Content,
	}, nil
}

func (m *mockMessageService) GetMessages(ctx context.Context) ([]domain.Message, error) {
	if m.hasError {
		return nil, errors.New("error")
	}

	return []domain.Message{
		{
			ID:      "1",
			Content: "Hello",
		},
		{
			ID:      "2",
			Content: "World",
		},
	}, nil
}
