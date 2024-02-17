package api

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"my-message-app/internal/domain"
	"my-message-app/internal/service"
)

const jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRlbW9AZXhhbXBsZS5jb20iLCJleHAiOjE3MDgxNjcyNTV9.wpvoLLQIkZMJbPdChaY2bN7JlnRF2Mr9Ln41aZZ9wyc"

func TestHandler_Register(t *testing.T) {
	type args struct {
		rw  http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "Register",
			args: args{
				rw: httptest.NewRecorder(),
				req: httptest.NewRequest(
					http.MethodPost,
					"/v1/register",
					strings.NewReader(`{"email":"test@t.c","password":"test"}`),
				),
			},
			want: &domain.User{
				Email: "test@t.c",
			},
		},
		{
			name: "RegisterError",
			args: args{
				rw: httptest.NewRecorder(),
				req: httptest.NewRequest(
					http.MethodPost,
					"/v1/register",
					strings.NewReader(`{"email":"test@t.c","password":"test"}`),
				),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newHandler(newServices(tt.wantErr))
			h.Register(tt.args.rw, tt.args.req)

			res := tt.args.rw.(*httptest.ResponseRecorder).Result()
			resBody, _ := io.ReadAll(res.Body)

			if !tt.wantErr && res.StatusCode != http.StatusCreated {
				t.Errorf("Handler.Register() status = %d, want %d", res.StatusCode, http.StatusCreated)
			}

			if tt.wantErr && res.StatusCode == http.StatusCreated {
				t.Errorf("Handler.Register() status = %d, want %d", res.StatusCode, http.StatusCreated)
			}

			if !tt.wantErr && !strings.Contains(string(resBody), tt.want.Email) {
				t.Errorf("Handler.Register() body = %s, want %s", string(resBody), tt.want.Email)
			}

			res.Body.Close()
		})
	}
}

func TestHandler_Login(t *testing.T) {
	type args struct {
		rw  http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Login",
			args: args{
				rw: httptest.NewRecorder(),
				req: httptest.NewRequest(
					http.MethodPost,
					"/v1/login",
					strings.NewReader(`{"email":"test@t.c","password":"test"}`),
				),
			},
			want: "__token__",
		},
		{
			name: "LoginError",
			args: args{
				rw: httptest.NewRecorder(),
				req: httptest.NewRequest(
					http.MethodPost,
					"/v1/login",
					strings.NewReader(`{"email":"test@t.c","password":"test"}`),
				),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newHandler(newServices(tt.wantErr))
			h.Login(tt.args.rw, tt.args.req)

			res := tt.args.rw.(*httptest.ResponseRecorder).Result()
			resBody, _ := io.ReadAll(res.Body)

			if !tt.wantErr && res.StatusCode != http.StatusOK {
				t.Errorf("Handler.Login() status = %d, want %d", res.StatusCode, http.StatusOK)
			}

			if tt.wantErr && res.StatusCode == http.StatusOK {
				t.Errorf("Handler.Login() status = %d, want %d", res.StatusCode, http.StatusOK)
			}

			if !tt.wantErr && !strings.Contains(string(resBody), tt.want) {
				t.Errorf("Handler.Login() body = %s, want %s", string(resBody), tt.want)
			}
		})
	}
}

func TestHandler_Logout(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/v1/logout",
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	type args struct {
		rw  http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Logout",
			args: args{
				rw:  httptest.NewRecorder(),
				req: req.Clone(context.Background()),
			},
			wantErr: false,
		},
		{
			name: "LogoutError",
			args: args{
				rw:  httptest.NewRecorder(),
				req: req.Clone(context.Background()),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newHandler(newServices(tt.wantErr))
			h.Logout(tt.args.rw, tt.args.req)

			res := tt.args.rw.(*httptest.ResponseRecorder).Result()

			if !tt.wantErr && res.StatusCode != http.StatusOK {
				t.Errorf("Handler.Logout() status = %d, want %d", res.StatusCode, http.StatusOK)
			}

			if tt.wantErr && res.StatusCode == http.StatusOK {
				t.Errorf("Handler.Logout() status = %d, want %d", res.StatusCode, http.StatusOK)
			}
		})
	}
}

func TestHandler_Authenticator(t *testing.T) {
	type args struct {
		authService service.AuthService
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Authenticator",
			args: args{
				authService: &mockAuthService{},
			},
			wantErr: false,
		},
		{
			name: "AuthenticatorError",
			args: args{
				authService: &mockAuthService{hasError: true},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newHandler(newServices(tt.wantErr))
			middleware := h.Authenticator(tt.args.authService)

			mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			httpHandler := middleware(mockHandler)

			rw := httptest.NewRecorder()
			req := httptest.NewRequest(
				http.MethodGet,
				"/v1/messages",
				nil,
			)
			req.Header.Set("Authorization", "Bearer "+jwtToken)

			httpHandler.ServeHTTP(rw, req)

			res := rw.Result()

			if !tt.wantErr && res.StatusCode != http.StatusOK {
				t.Errorf("Handler.Authenticator() status = %d, want %d", res.StatusCode, http.StatusOK)
			}

			if tt.wantErr && res.StatusCode == http.StatusOK {
				t.Errorf("Handler.Authenticator() status = %d, want %d", res.StatusCode, http.StatusOK)
			}
		})
	}
}

func TestFindToken(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/v1/messages",
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FindToken",
			args: args{
				r: req,
			},
			want: jwtToken,
		},
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
	req := httptest.NewRequest(
		http.MethodGet,
		"/v1/messages",
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TokenFromHeader",
			args: args{
				r: req,
			},
			want: jwtToken,
		},
		{
			name: "TokenFromHeaderNoToken",
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
		{
			name: "TokenFromQuery",
			args: args{
				r: httptest.NewRequest(
					http.MethodGet,
					"/v1/messages?jwt="+jwtToken,
					nil,
				),
			},
			want: jwtToken,
		},
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
	req := httptest.NewRequest(
		http.MethodGet,
		"/v1/messages",
		nil,
	)

	cookie := &http.Cookie{
		Name:  "jwt",
		Value: jwtToken,
	}

	req.AddCookie(cookie)

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TokenFromCookie",
			args: args{
				r: req,
			},
			want: jwtToken,
		},
		{
			name: "TokenFromCookieNoCookie",
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
	return "__token__", nil
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
