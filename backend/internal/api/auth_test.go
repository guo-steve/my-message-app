package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"my-message-app/internal/domain"
	"my-message-app/internal/service"

	"github.com/go-chi/httplog/v2"
)

func TestHandler_Login(t *testing.T) {
	type fields struct {
		services *service.Services
		logger   *httplog.Logger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				services: tt.fields.services,
				logger:   tt.fields.logger,
			}
			h.Login(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_Authenticator(t *testing.T) {
	type fields struct {
		services *service.Services
		logger   *httplog.Logger
	}
	type args struct {
		authService service.AuthService
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   func(http.Handler) http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				services: tt.fields.services,
				logger:   tt.fields.logger,
			}
			if got := h.Authenticator(tt.args.authService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.Authenticator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindToken(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindToken(tt.args.r); got != tt.want {
				t.Errorf("FindToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenFromHeader(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TokenFromHeader(tt.args.r); got != tt.want {
				t.Errorf("TokenFromHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenFromQuery(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TokenFromQuery(tt.args.r); got != tt.want {
				t.Errorf("TokenFromQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenFromCookie(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestTokenFromCookie",
			args: args{
				r: httptest.NewRequest(
					http.MethodGet,
					"/v1/messages",
					nil,
				),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TokenFromCookie(tt.args.r); got != tt.want {
				t.Errorf("TokenFromCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}

var _ service.AuthService = &mockAuthService{}

type mockAuthService struct {
	hasError bool
}

func (m *mockAuthService) Login(ctx context.Context, creds domain.Credentials) (string, error) {
	if m.hasError {
		return "", errors.New("error")
	}
	return "", nil
}

func (m *mockAuthService) Authenticate(ctx context.Context, tokenString string) (*domain.User, error) {
	if m.hasError {
		return nil, errors.New("error")
	}
	return &domain.User{
		ID: "1",
	}, nil
}

func (m *mockAuthService) Register(ctx context.Context, user domain.User) (*domain.User, error) {
	if m.hasError {
		return nil, errors.New("error")
	}
	return &user, nil
}

func (m *mockAuthService) Logout(ctx context.Context, tokenString string) error {
	if m.hasError {
		return errors.New("error")
	}
	return nil
}
