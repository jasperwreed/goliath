package interfaces

import "github.com/jasperwreed/goliath/pkg/models"

type UserRepository interface {
	FindByID(id uint) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByPhoneNumber(phoneNumber string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(user *models.User) error
}
