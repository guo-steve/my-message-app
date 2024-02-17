package api

import (
	"net/http"

	"my-message-app/internal/service"

	"github.com/go-chi/httplog/v2"
)

// Handler is a struct that holds the repository
type Handler struct {
	services *service.Services
	logger   *httplog.Logger
}

// NewHandler returns a new Handler
func NewHandler(services *service.Services, logger *httplog.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

// Ping
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
