package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mauroccvieira/restapi-echo-go/logger"
	"github.com/mauroccvieira/restapi-echo-go/models"
	"github.com/mauroccvieira/restapi-echo-go/services"

	"github.com/stretchr/testify/assert"
)

type MockUserService struct {
	services.UserService
	MockGetUsers       func() ([]models.User, error)
	MockDeleteUserById func(id int) error
}

func (m *MockUserService) GetUsers() ([]models.User, error) {
	return m.MockGetUsers()
}

func (m *MockUserService) DeleteUser(id int) error {
	return m.MockDeleteUserById(id)
}

func setUp_user_test() func() {
	logger.New()

	return func() {
		logger.Sync()
		logger.Delete()
	}
}

func TestGetUsersSuccessCase(t *testing.T) {
	defer setUp_user_test()()

	s := &MockUserService{
		MockGetUsers: func() ([]models.User, error) {
			var r []models.User
			return r, nil
		},
	}

	mockService := &services.Services{User: s}

	e := Echo()
	h := New(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.UserHandler.GetUsers(c))
	assert.Equal(t, rec.Code, http.StatusOK)
}

func TestGetUsers500Case(t *testing.T) {
	defer setUp_user_test()()

	s := &MockUserService{
		MockGetUsers: func() ([]models.User, error) {
			var r []models.User
			return r, errors.New("fake error")
		},
	}

	mockService := &services.Services{User: s}

	e := Echo()
	h := New(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.UserHandler.GetUsers(c))
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
