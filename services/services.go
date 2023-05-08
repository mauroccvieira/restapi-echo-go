package services

import "github.com/mauroccvieira/restapi-echo-go/stores"

type Services struct {
	User UserService
}

func New(s *stores.Stores) *Services {
	return &Services{
		User: &userService{stores: s},
	}
}
