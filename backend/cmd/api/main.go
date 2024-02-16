package main

import (
	"database/sql"
	"net/http"
	"os"

	"my-message-app/internal/api"
	"my-message-app/internal/repo/sqlite"
	"my-message-app/internal/service"

	_ "github.com/glebarez/go-sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog/v2"
)

func main() {
	// Logger
	logger := httplog.NewLogger("my-message-app", httplog.Options{
		JSON:            true,
		Concise:         true,
		SourceFieldName: "caller",
	})

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	}))

	// Connect to the database
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		logger.Error("failed to open database", err)
		os.Exit(1)
	}

	// Create the messages table
	// This is usually done in a migration outside of the application
	_, err = db.Exec(`CREATE TABLE messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		logger.Error("failed to create table", err)
		os.Exit(1)
	}

	repo := sqlite.NewSqliteRepo(db)
	services := service.NewServices(service.NewMessageService(repo))
	handler := api.NewHandler(services, logger.With("component", "api"))

	r.Route("/v1/messages", func(r chi.Router) {
		r.Post("/", handler.PostMessage) // POST /messages
		r.Get("/", handler.GetMessages)  // GET /messages
	})

	logger.Info("server started on :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("server failed", err)
		os.Exit(1)
	}
}
