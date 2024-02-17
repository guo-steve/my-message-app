package api

import (
	"encoding/json"
	"net/http"

	"my-message-app/internal/domain"
)

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

	user, ok := req.Context().Value("user").(domain.User)
	if !ok {
		http.Error(resWtr, "user not found", http.StatusUnauthorized)
		return
	}

	message := domain.Message{
		Content:   postMessageRequest.Content,
		CreatedBy: user.ID,
	}

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
