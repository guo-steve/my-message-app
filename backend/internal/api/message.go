package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"my-message-app/internal/domain"
	"my-message-app/internal/service"
)

// Handler is a struct that holds the repository
type Handler struct {
	services *service.Services
	logger   *slog.Logger
}

// NewHandler returns a new Handler
func NewHandler(services *service.Services, logger *slog.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

type PostMessageRequest struct {
	Content string `json:"content"`
}

// PostMessage handles the POST /messages endpoint
func (h *Handler) PostMessage(resWtr http.ResponseWriter, req *http.Request) {
	var postMessageRequest PostMessageRequest

	if err := json.NewDecoder(req.Body).Decode(&postMessageRequest); err != nil {
		h.logger.Error(err.Error())
		http.Error(resWtr, err.Error(), http.StatusBadRequest)
		return
	}

	message := domain.Message{Content: postMessageRequest.Content}

	result, err := h.services.MessageService.CreateMessage(req.Context(), message)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(resWtr, err.Error(), http.StatusInternalServerError)
		return
	}

	resWtr.WriteHeader(http.StatusCreated)
	json.NewEncoder(resWtr).Encode(result)
}

// GetMessages handles the GET /messages endpoint
func (h *Handler) GetMessages(resWtr http.ResponseWriter, req *http.Request) {
	messages, err := h.services.MessageService.GetMessages(req.Context())
	if err != nil {
		http.Error(resWtr, err.Error(), http.StatusInternalServerError)
		return
	}

	resWtr.WriteHeader(http.StatusOK)
	json.NewEncoder(resWtr).Encode(messages)
}
