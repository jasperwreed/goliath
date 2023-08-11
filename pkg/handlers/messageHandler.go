package handlers

import (
	"net/http"

	"github.com/jasperwreed/goliath/pkg/services"
)

type MessageHandler struct {
	service *services.MessageService
}

func NewMessageHandler(s *services.MessageService) *MessageHandler {
	return &MessageHandler{service: s}
}

func (h *MessageHandler) GetMessageByID(w http.ResponseWriter, r *http.Request) {
	// Handle the request and response here...
}

func (h *MessageHandler) GetMessagesByConversationID(w http.ResponseWriter, r *http.Request) {
	// Handle the request and response here...
}
