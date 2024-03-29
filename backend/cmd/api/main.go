package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"my-message-app/internal/api"
	"my-message-app/internal/repo/sqlite"
	"my-message-app/internal/service"

	_ "github.com/glebarez/go-sqlite"
	"github.com/go-chi/httplog/v2"
)

func main() {
	// Logger
	logger := httplog.NewLogger("my-message-app", httplog.Options{
		JSON:            true,
		Concise:         true,
		SourceFieldName: "caller",
		ReplaceAttrsOverride: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "caller" {
				source := a.Value.Any().(*slog.Source)
				a.Value = slog.StringValue(fmt.Sprintf("%s:%d", source.File, source.Line))
				return slog.Attr{Key: "caller", Value: a.Value}
			}
			return a
		},
	})

	// Connect to the database
	// db, err := sql.Open("sqlite", ":memory:")
	db, err := sql.Open("sqlite", "./dev.db")
	if err != nil {
		logger.Error("failed to open database", err)
		os.Exit(1)
	}

	// Create the messages table
	// This is usually done in a migration outside of the application
	sqls, err := os.ReadFile("./database/schema.sql")
	if err != nil {
		logger.Error("failed to read schema.sql", err)
		os.Exit(1)
	}
	_, err = db.Exec(string(sqls))
	if err != nil {
		logger.Error("failed to create table", err)
	}

	repo := sqlite.NewSqliteRepo(db)
	services := service.NewServices(service.NewMessageService(repo), service.NewAuthService(repo, repo), service.NewUserService(repo))
	handler := api.NewHandler(services, logger)
	r := api.InitRounter(handler)

	logger.Info("server started on :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("server failed", err)
		os.Exit(1)
	}
}
