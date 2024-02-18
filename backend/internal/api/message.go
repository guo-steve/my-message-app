package api

import (
	"context"
	"encoding/json"
	"net/http"

	"my-message-app/internal/domain"
)

type PostMessageRequest struct {
	Content string `json:"content"`
}

func (h *Handler) replaceCreatedBy(ctx context.Context, message domain.Message, user domain.User) (*domain.Message, error) {
	if message.CreatedBy == user.ID {
		message.CreatedBy = "You"
	} else {
		usr, err := h.services.UserService.GetUserByID(ctx, message.CreatedBy)
		if err != nil {
			return nil, err
		}
		message.CreatedBy = usr.FullName
		if message.CreatedBy == "" {
			message.CreatedBy = usr.Email
		}
	}
	return &message, nil
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

	result, err = h.replaceCreatedBy(req.Context(), *result, user)
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
	user, ok := req.Context().Value("user").(domain.User)
	if !ok {
		http.Error(resWtr, "user not found", http.StatusUnauthorized)
		return
	}

	messages, err := h.services.MessageService.GetMessages(req.Context())
	if err != nil {
		http.Error(resWtr, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := range messages {
		newMsg, err := h.replaceCreatedBy(req.Context(), messages[i], user)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(resWtr, err.Error(), http.StatusInternalServerError)
			return
		}
		messages[i] = *newMsg
	}

	resWtr.WriteHeader(http.StatusOK)
	json.NewEncoder(resWtr).Encode(messages)
}
