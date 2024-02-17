package main

import (
	"database/sql"
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
	})

	// Connect to the database
	db, err := sql.Open("sqlite", ":memory:")
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
		os.Exit(1)
	}

	// tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("MMA_JWT_KEY")), nil)

	repo := sqlite.NewSqliteRepo(db)
	services := service.NewServices(service.NewMessageService(repo), service.NewAuthService(repo, repo))
	handler := api.NewHandler(services, logger)
	r := api.InitRounter(handler)

	logger.Info("server started on :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("server failed", err)
		os.Exit(1)
	}
}
