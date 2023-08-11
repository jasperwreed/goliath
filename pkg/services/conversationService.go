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
