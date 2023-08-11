package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jasperwreed/goliath/pkg/services"
)

type MessageHandler struct {
	service *services.MessageService
}

func NewMessageHandler(s *services.MessageService) *MessageHandler {
	return &MessageHandler{service: s}
}

// Create a new message
func (h *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var newMsg models.Message
	err := json.NewDecoder(r.Body).Decode(&newMsg)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	err = h.service.CreateMessage(&newMsg)
	if err != nil {
		http.Error(w, "Failed to create message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMsg)
}

// Get a message by ID
func (h *MessageHandler) GetMessageByID(w http.ResponseWriter, r *http.Request) {
	strID := /* extract from request */

	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	message, err := h.service.GetMessageByID(uint(id))
	if err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(message)
}

// Get all messages for a specific conversation
func (h *MessageHandler) GetMessagesByConversationID(w http.ResponseWriter, r *http.Request) {
	strConversationID := /* extract conversation ID from request */

	conversationID, err := strconv.Atoi(strConversationID)
	if err != nil {
		http.Error(w, "Invalid conversation ID format", http.StatusBadRequest)
		return
	}

	messages, err := h.service.GetMessagesByConversationID(uint(conversationID))
	if err != nil {
		http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(messages)
}

// Delete a message
func (h *MessageHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	strID := /* extract from request */

	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteMessage(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content response
}
