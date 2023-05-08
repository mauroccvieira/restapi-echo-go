package services

import (
	"github.com/mauroccvieira/restapi-echo-go/models"
	"github.com/mauroccvieira/restapi-echo-go/stores"
)

type (
	UserService interface {
		GetUsers() ([]models.User, error)
		CreateUser(user *models.User) (models.User, error)
		// GetUser(id int) (models.User, error)
		// UpdateUser(user models.User) (models.User, error)
		// DeleteUser(id int) (models.User, error)
	}

	userService struct {
		stores *stores.Stores
	}
)

func (s *userService) GetUsers() ([]models.User, error) {
	r, err := s.stores.User.Get(nil)
	return r, err
}

func (s *userService) CreateUser(user *models.User) (models.User, error) {
	r, err := s.stores.User.Create(nil, user)
	return r, err
}
