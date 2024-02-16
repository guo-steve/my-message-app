package sqlite

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"my-message-app/internal/domain"

	_ "github.com/glebarez/go-sqlite"
)

func getTestDB() *sql.DB {
	db, _ := sql.Open("sqlite", ":memory:")
	_, _ = db.Exec(`CREATE TABLE messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	return db
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
		// {
		// 	name: "PostMessage error",
		// 	fields: fields{
		// 		db: testDb,
		// 	},
		// 	args: args{
		// 		ctx: context.Background(),
		// 		message: domain.Message{
		// 			Content: "Hello, World!",
		// 		},
		// 	},
		// 	want: &domain.Message{
		// 		Content: "Hello, World!",
		// 	},
		// 	wantErr: true,
		// },
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
		ctx context.Context
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
				ctx: context.Background(),
			},
			want: []domain.Message{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SqliteRepo{
				db: tt.fields.db,
			}
			_, err := r.GetMessages(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SqliteRepo.GetMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
