package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"my-message-app/internal/domain"
	"my-message-app/internal/repo"
)

type SqliteRepo struct {
	db *sql.DB
}

var (
	_ repo.MessageRepo      = &SqliteRepo{}
	_ repo.UserRepo         = &SqliteRepo{}
	_ repo.InvalidTokenRepo = &SqliteRepo{}
)

func NewSqliteRepo(db *sql.DB) *SqliteRepo {
	return &SqliteRepo{db: db}
}

// CreateMessage adds a message to the database
func (r *SqliteRepo) CreateMessage(ctx context.Context, message domain.Message) (*domain.Message, error) {
	row := r.db.QueryRowContext(
		ctx,
		"INSERT INTO messages (content, created_by) VALUES (?, ?) RETURNING *",
		message.Content,
		message.CreatedBy,
	)
	err := row.Scan(&message.ID, &message.Content, &message.CreatedBy, &message.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to insert message: %w", err)
	}
	return &message, err
}

// GetMessages returns all messages from the database
func (r *SqliteRepo) GetMessages(ctx context.Context, createdBy string) ([]domain.Message, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, content, created_by, created_at
		FROM messages
		WHERE created_by = ?`,
		createdBy,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}
	defer rows.Close()

	messages := make([]domain.Message, 0)

	for rows.Next() {
		var message domain.Message
		if err := rows.Scan(&message.ID, &message.Content, &message.CreatedBy, &message.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

// FindUserByEmail returns a user by email
func (r *SqliteRepo) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT id, email, password FROM users
		WHERE email = ? AND active = 1`,
		email,
	)
	var user domain.User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return &user, nil
}

// CreateUser adds a user to the database
func (r *SqliteRepo) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	row := r.db.QueryRowContext(
		ctx,
		`INSERT INTO users (email, password, full_name) VALUES (?, ?, ?)
		RETURNING
		id, email, full_name, active, created_at, updated_at`,
		user.Email,
		user.Password,
		user.FullName,
	)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FullName,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}
	return &user, err
}

// CreateToken adds a token to the database
func (r *SqliteRepo) CreateInvalidToken(ctx context.Context, token string) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO invalid_tokens (token) VALUES (?)",
		token,
	)
	if err != nil {
		return fmt.Errorf("failed to insert token: %w", err)
	}
	return err
}

// FindToken returns a token by token string
func (r *SqliteRepo) FindInvalidToken(ctx context.Context, tokenString string) (*domain.InvalidToken, error) {
	row := r.db.QueryRowContext(
		ctx,
		"SELECT token, created_at FROM invalid_tokens WHERE token = ?",
		tokenString,
	)

	var token domain.InvalidToken
	err := row.Scan(&token.Token, &token.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find token: %w", err)
	}
	return &token, nil
}

// GetUserByID returns a user by id
func (r *SqliteRepo) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	row := r.db.QueryRowContext(
		ctx,
		`SELECT id, email, full_name, active, created_at, updated_at
		FROM users WHERE id = ?`,
		id,
	)

	var user domain.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FullName,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return &user, nil
}
