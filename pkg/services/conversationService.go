package services

import (
	"github.com/jasperwreed/goliath/interfaces"
	"github.com/jasperwreed/goliath/models"
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
