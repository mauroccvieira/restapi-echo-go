package services

import (
	"github.com/mauroccvieira/restapi-echo-go/models"
	"github.com/mauroccvieira/restapi-echo-go/stores"
)

type (
	UserService interface {
		GetUsers() ([]models.User, error)
		// GetUser(id int) (models.User, error)
		// CreateUser(user models.User) (models.User, error)
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

// func (s *userService) GetUser(id int) (models.User, error) {
// 	return models.User{}, nil
// }

// func (s *userService) CreateUser(user models.User) (models.User, error) {
// 	return models.User{}, nil
// }

// func (s *userService) UpdateUser(user models.User) (models.User, error) {
// 	return models.User{}, nil
// }

// func (s *userService) DeleteUser(id int) (models.User, error) {
// 	return models.User{}, nil
// }
