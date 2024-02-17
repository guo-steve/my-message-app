package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"my-message-app/internal/domain"
	"my-message-app/internal/service"
)

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (h *Handler) Register(rw http.ResponseWriter, req *http.Request) {
	// Decode the request body into domain.User
	var user domain.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the Register method from the AuthService
	savedUser, err := h.services.AuthService.Register(req.Context(), user)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// Write the token string to the response
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(UserResponse{
		ID:        savedUser.ID,
		Email:     savedUser.Email,
		Fullname:  savedUser.FullName,
		Active:    savedUser.Active,
		CreatedAt: savedUser.CreatedAt.String(),
		UpdatedAt: savedUser.UpdatedAt.String(),
	})
}

func (h *Handler) Login(rw http.ResponseWriter, req *http.Request) {
	// Decode the request body into domain.Credentials
	var creds domain.Credentials
	if err := json.NewDecoder(req.Body).Decode(&creds); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the Login method from the AuthService
	tokenString, err := h.services.AuthService.Login(req.Context(), creds)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusUnauthorized)
		return
	}

	// Write the token string to the response
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]string{"token": tokenString})
}

func (h *Handler) Logout(rw http.ResponseWriter, req *http.Request) {
	// Get token from request
	tokenString := FindToken(req)
	if tokenString == "" {
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Call the Logout method from the AuthService
	err := h.services.AuthService.Logout(req.Context(), tokenString)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusUnauthorized)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (h *Handler) Authenticator(authService service.AuthService) func(http.Handler) http.Handler {
	// return middleware
	return func(next http.Handler) http.Handler {
		// return HandlerFunc
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := FindToken(r)
			if tokenString == "" {
				h.logger.Error(fmt.Sprintf("Cannot find token"))
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := r.Context()

			user, err := authService.Authenticate(ctx, tokenString)
			if err != nil {
				h.logger.Error(fmt.Sprintf("Authenticate err: %v", err))
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// create new context with user value
			newCtx := context.WithValue(ctx, "user", *user)

			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}

func FindToken(r *http.Request) string {
	for _, f := range []func(*http.Request) string{
		TokenFromQuery,
		TokenFromHeader,
		TokenFromCookie,
	} {
		if token := f(r); token != "" {
			return token
		}
	}

	return ""
}

// TokenFromHeader tries to retreive the token string from the
// "Authorization" reqeust header: "Authorization: BEARER T".
func TokenFromHeader(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}

// TokenFromQuery tries to retreive the token string from the "jwt" URI
// query parameter.
//
// To use it, build our own middleware handler, such as:
//
//	func Verifier(ja *JWTAuth) func(http.Handler) http.Handler {
//		return func(next http.Handler) http.Handler {
//			return Verify(ja, TokenFromQuery, TokenFromHeader, TokenFromCookie)(next)
//		}
//	}
func TokenFromQuery(r *http.Request) string {
	// Get token from query param named "jwt".
	return r.URL.Query().Get("jwt")
}

// TokenFromCookie tries to retreive the token string from a cookie named
// "jwt".
func TokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return ""
	}
	return cookie.Value
}
