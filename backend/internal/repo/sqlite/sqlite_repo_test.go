package sqlite

import (
	"context"
	"database/sql"
	"os"
	"reflect"
	"testing"

	"my-message-app/internal/domain"

	_ "github.com/glebarez/go-sqlite"
)

func getTestDB() *sql.DB {
	db, _ := sql.Open("sqlite", ":memory:")
	sqls, _ := os.ReadFile("../../../database/schema.sql")
	db.Exec(string(sqls))
	return db
}

func createTestUser(db *sql.DB) string {
	row := db.QueryRow(`INSERT INTO users (email, password)
		VALUES (tester@example.com', 'password')
		RETURNING id`)

	var id string
	row.Scan(&id)
	return id
}

func createTestMessage(db *sql.DB, createdBy string) *domain.Message {
	row := db.QueryRow(`INSERT INTO messages (content, created_by)
		VALUES ('Hello, World!', ?)
		RETURNING *`, createdBy)

	var message domain.Message
	row.Scan(&message.ID, &message.Content, &message.CreatedBy, &message.CreatedAt)
	return &message
}

func TestNewSqliteRepo(t *testing.T) {
	testDB := getTestDB()

	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *SqliteRepo
	}{
		{
			name: "NewSqliteRepo",
			args: args{
				db: testDB,
			},
			want: &SqliteRepo{
				db: testDB,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSqliteRepo(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSqliteRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqliteRepo_PostMessage(t *testing.T) {
	testDb := getTestDB()

	type fields struct {
		db *sql.DB
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
			name: "PostMessage",
			fields: fields{
				db: testDb,
			},
			args: args{
				ctx: context.Background(),
				message: domain.Message{
					Content: "Hello, World!",
				},
			},
			want: &domain.Message{
				Content: "Hello, World!",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SqliteRepo{
				db: tt.fields.db,
			}
			got, err := r.CreateMessage(tt.args.ctx, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("SqliteRepo.PostMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Content != tt.want.Content {
				t.Errorf("SqliteRepo.PostMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqliteRepo_GetMessages(t *testing.T) {
	testDb := getTestDB()

	type fields struct {
		db *sql.DB
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
			name: "GetMessages",
			fields: fields{
				db: testDb,
			},
			args: args{
				ctx:       context.Background(),
				createdBy: "1",
			},
			want: []domain.Message{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SqliteRepo{
				db: tt.fields.db,
			}
			_, err := r.GetMessages(tt.args.ctx, tt.args.createdBy)
			if (err != nil) != tt.wantErr {
				t.Errorf("SqliteRepo.GetMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSqliteRepo_UpdateMessage(t *testing.T) {
	testDb := getTestDB()
	userID := createTestUser(testDb)
	testMessage := createTestMessage(testDb, userID)

	testMessage.Content = "Updated content"

	type fields struct {
		db *sql.DB
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
			name: "UpdateMessage",
			fields: fields{
				db: testDb,
			},
			args: args{
				ctx:     context.Background(),
				message: *testMessage,
			},
			want: &domain.Message{
				ID:        testMessage.ID,
				Content:   "Updated content",
				CreatedBy: testMessage.CreatedBy,
				CreatedAt: testMessage.CreatedAt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SqliteRepo{
				db: tt.fields.db,
			}
			got, err := r.UpdateMessage(tt.args.ctx, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("SqliteRepo.UpdateMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Content != tt.want.Content {
				t.Errorf("SqliteRepo.UpdateMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqliteRepo_DeleteMessage(t *testing.T) {
	testDb := getTestDB()

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx       context.Context
		messageID string
		userID    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "DeleteMessage",
			fields: fields{
				db: testDb,
			},
			args: args{
				ctx:       context.Background(),
				messageID: "1",
				userID:    "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SqliteRepo{
				db: tt.fields.db,
			}
			if err := r.DeleteMessage(tt.args.ctx, tt.args.messageID, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SqliteRepo.DeleteMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
