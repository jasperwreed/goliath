package interfaces

import (
	"github.com/jasperwreed/goliath/models"
)

type MessageRepository interface {
	FindByID(id string) (*models.Message, error)
	FindByConversationID(conversationID string) ([]models.Message, error)
	Save(message *models.Message) error
	Delete(message *models.Message) error
}
