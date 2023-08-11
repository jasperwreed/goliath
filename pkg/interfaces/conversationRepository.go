package interfaces

import (
	"github.com/jasperwreed/goliath/pkg/models"
)

type ConversationRepository interface {
	FindByID(id uint) (*models.Conversation, error)
	Create(conversation *models.Conversation) error
	Update(conversation *models.Conversation) error
	Delete(conversation *models.Conversation) error
}
