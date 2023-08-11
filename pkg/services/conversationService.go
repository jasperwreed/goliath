package services

import (
	"github.com/jasperwreed/goliath/pkg/interfaces"
	"github.com/jasperwreed/goliath/pkg/models"
)

type ConversationService struct {
	repo interfaces.ConversationRepository
}

func NewConversationService(r interfaces.ConversationRepository) *ConversationService {
	return &ConversationService{repo: r}
}

func (s *ConversationService) GetConversationByID(id uint) (*models.Conversation, error) {
	return s.repo.FindByID(id)
}

func (s *ConversationService) CreateConversation(conversation *models.Conversation) error {
	return s.repo.Create(conversation)
}

func (s *ConversationService) UpdateConversation(conversation *models.Conversation) error {
	return s.repo.Update(conversation)
}

func (s *ConversationService) DeleteConversation(id uint) error {
	conversation, err := s.GetConversationByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(conversation)
}
