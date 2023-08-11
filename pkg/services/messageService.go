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

func (s *MessageService) CreateMessage(message *models.Message) error {
	return s.repo.Create(message)
}

func (s *MessageService) DeleteMessage(id string) error {
	message, err := s.GetMessageByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(message)
}

func (s *MessageService) GetMessagesByConversationID(conversationID string) ([]models.Message, error) {
	return s.repo.FindByConversationID(conversationID)
}
