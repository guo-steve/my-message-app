package api

import (
	"encoding/json"
	"net/http"

	"my-message-app/internal/domain"
	"my-message-app/internal/repo"
)

// Handler is a struct that holds the repository
type Handler struct {
	repo repo.Repository
}

// NewHandler returns a new Handler
func NewHandler(repo repo.Repository) *Handler {
	return &Handler{repo: repo}
}

type PostMessageRequest struct {
	Content string `json:"content"`
}

// PostMessage handles the POST /messages endpoint
func (h *Handler) PostMessage(resWtr http.ResponseWriter, req *http.Request) {
	var postMessageRequest PostMessageRequest

	if err := json.NewDecoder(req.Body).Decode(&postMessageRequest); err != nil {
		http.Error(resWtr, err.Error(), http.StatusBadRequest)
		return
	}

	message := domain.Message{Content: postMessageRequest.Content}

	result, err := h.repo.PostMessage(req.Context(), message)
	if err != nil {
		http.Error(resWtr, err.Error(), http.StatusInternalServerError)
		return
	}

	resWtr.WriteHeader(http.StatusCreated)
	json.NewEncoder(resWtr).Encode(result)
}

// GetMessages handles the GET /messages endpoint
func (h *Handler) GetMessages(resWtr http.ResponseWriter, req *http.Request) {
	messages, err := h.repo.GetMessages(req.Context())
	if err != nil {
		http.Error(resWtr, err.Error(), http.StatusInternalServerError)
		return
	}

	resWtr.WriteHeader(http.StatusOK)
	json.NewEncoder(resWtr).Encode(messages)
}
