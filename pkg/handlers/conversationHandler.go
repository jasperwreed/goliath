package handlers

import (
	"net/http"
	"github.com/jasperwreed/goliath/pkg/services"
)

type ConversationHandler struct {
	service *services.ConversationService
}

func NewConversationHandler(s *services.ConversationService) *ConversationHandler {
	return &ConversationHandler{service: s}
}

func (h *ConversationHandler) GetConversationByID(w http.ResponseWriter, r *http.Request) {
	// Handle the request and response here...
}
