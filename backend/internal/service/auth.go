package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"my-message-app/internal/domain"
	"my-message-app/internal/repo"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtKey = []byte(os.Getenv("MMA_JWT_KEY"))

	ErrInvalidToken = fmt.Errorf("invalid token")
)

func init() {
	if len(jwtKey) == 0 {
		jwtKey = []byte("a4aWevITOb6NTVitH5xrfGJfpVBcmBcZ")
	}
}

type authService struct {
	userRepo  repo.UserRepo
	tokenRepo repo.InvalidTokenRepo
}

func NewAuthService(userRepo repo.UserRepo, tokenRepo repo.InvalidTokenRepo) *authService {
	return &authService{userRepo: userRepo, tokenRepo: tokenRepo}
}

func (a *authService) Register(ctx context.Context, user domain.User) (*domain.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password error %w", err)
	}

	user.Password = string(hash)

	savedUser, errC := a.userRepo.CreateUser(ctx, user)
	if errC != nil {
		return nil, fmt.Errorf("create user error %w", errC)
	}

	return savedUser, nil
}

func (a *authService) Login(ctx context.Context, creds domain.Credentials) (string, error) {
	user, err := a.userRepo.FindUserByEmail(ctx, creds.Email)
	if err != nil {
		return "", fmt.Errorf("find user error %w", err)
	}

	errBc := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if errBc != nil {
		return "", fmt.Errorf("invalid password %w", errBc)
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &domain.JWTClaims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("sign token error %w", err)
	}

	return tokenString, nil
}

func (a *authService) Authenticate(ctx context.Context, tokenString string) (*domain.User, error) {
	claims := &domain.JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse token error %w", err)
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	user, err := a.userRepo.FindUserByEmail(ctx, claims.Email)
	if err != nil {
		return nil, fmt.Errorf("find user error %w", err)
	}

	invalidToken, errT := a.tokenRepo.FindInvalidToken(ctx, tokenString)
	if errT != nil {
		return nil, fmt.Errorf("find token error %w", errT)
	}

	if invalidToken != nil {
		return nil, ErrInvalidToken
	}

	return user, nil
}

func (a *authService) Logout(ctx context.Context, tokenString string) error {
	err := a.tokenRepo.CreateInvalidToken(ctx, tokenString)
	if err != nil {
		return fmt.Errorf("invalidate token error %w", err)
	}

	return nil
}
