package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FullName  string    `json:"full_name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type Message struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

type RichMessage struct {
	Message
	CreatedBy User `json:"created_by"`
}

type InvalidToken struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}
