package helpers

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jasperwreed/goliath/pkg/models"
	"github.com/jasperwreed/goliath/pkg/services"
)

/*
turn this into a helper function
id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	conversation, err := h.service.GetConversationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conversation not found"})
		return
	}
*/

// HandlerHelper is a helper function for handlers.
type HandlerHelper struct {
	service *services.ConversationService
}

// NewHandlerHelper creates a new HandlerHelper.
func NewHandlerHelper(s *services.ConversationService) *HandlerHelper {
	return &HandlerHelper{service: s}
}


func (h *HandlerHelper) handlerHelper(strID, c *gin.Context) {
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	conversation, err := h.service.GetConversationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conversation not found"})
		return
	}
}

