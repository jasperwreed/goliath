package services

import (
	"github.com/jasperwreed/goliath/pkg/models"
	"github.com/jasperwreed/goliath/pkg/interfaces"
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(r interfaces.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetByUsername(username string) (*models.User, error) {
	return s.repo.FindByUsername(username)
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) GetByPhoneNumber(phoneNumber string) (*models.User, error) {
	return s.repo.FindByPhoneNumber(phoneNumber)
}

func (s *UserService) Create(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) Update(user *models.User) error {
	return s.repo.Update(user)
}

func (s *UserService) Delete(user *models.User) error {
	return s.repo.Delete(user)
}
