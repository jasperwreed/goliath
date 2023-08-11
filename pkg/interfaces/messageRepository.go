package interfaces

import (
	"github.com/jasperwreed/goliath/pkg/models"
)

type MessageRepository interface {
	FindByID(id string) (*models.Message, error)
	FindByConversationID(conversationID string) ([]models.Message, error)
	Create(message *models.Message) error
	Delete(message *models.Message) error
}
