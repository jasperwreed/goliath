package handlers

import (
	"encoding/json"
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
	strID := /* extract from request */

	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	conversation, err := h.service.GetConversationByID(uint(id))
	if err != nil {
		http.Error(w, "Conversation not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(conversation)
}

func (h *ConversationHandler) CreateConversation(w http.ResponseWriter, r *http.Request) {
	var newConvo models.Conversation

	err := json.NewDecoder(r.Body).Decode(&newConvo)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	err = h.service.CreateConversation(&newConvo)
	if err != nil {
		http.Error(w, "Failed to create conversation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newConvo)
}

func (h *ConversationHandler) UpdateConversation(w http.ResponseWriter, r *http.Request) {
	strID := /* extract from request */

	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedConvo models.Conversation
	err = json.NewDecoder(r.Body).Decode(&updatedConvo)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	updatedConvo.ID = uint(id)
	err = h.service.UpdateConversation(&updatedConvo)
	if err != nil {
		http.Error(w, "Failed to update conversation", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedConvo)
}

func (h *ConversationHandler) DeleteConversation(w http.ResponseWriter, r *http.Request) {
	strID := /* extract from request */

	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteConversation(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete conversation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content response
}
