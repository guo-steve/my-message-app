package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"my-message-app/internal/domain"
)

type SqliteRepo struct {
	db *sql.DB
}

func NewSqliteRepo(db *sql.DB) *SqliteRepo {
	return &SqliteRepo{db: db}
}

// PostMessage adds a message to the database
func (r *SqliteRepo) PostMessage(ctx context.Context, message domain.Message) (*domain.Message, error) {
	row := r.db.QueryRowContext(
		ctx,
		"INSERT INTO messages (content) VALUES (?) RETURNING *",
		message.Content,
	)
	err := row.Scan(&message.ID, &message.Content, &message.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to insert message: %w", err)
	}
	return &message, err
}

// GetMessages returns all messages from the database
func (r *SqliteRepo) GetMessages(ctx context.Context) ([]domain.Message, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, content, created_at FROM messages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := make([]domain.Message, 0)

	for rows.Next() {
		var message domain.Message
		if err := rows.Scan(&message.ID, &message.Content, &message.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
