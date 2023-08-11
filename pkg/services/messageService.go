package services

import (
	"github.com/jasperwreed/goliath/pkg/interfaces"
	"github.com/jasperwreed/goliath/pkg/models"
)

type MessageService struct {
	repo interfaces.MessageRepository
}

func NewMessageService(r interfaces.MessageRepository) *MessageService {
	return &MessageService{repo: r}
}

func (s *MessageService) GetMessageByID(id string) (*models.Message, error) {
	return s.repo.FindByID(id)
}

func (s *MessageService) GetMessagesByConversationID(conversationID string) ([]models.Message, error) {
	return s.repo.FindByConversationID(conversationID)
}
