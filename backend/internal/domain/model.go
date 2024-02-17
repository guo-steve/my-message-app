package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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
	jwt.StandardClaims
}

type Message struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

type InvalidToken struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}
