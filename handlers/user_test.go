package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/mauroccvieira/restapi-echo-go/logger"
	"github.com/mauroccvieira/restapi-echo-go/models"
	"github.com/mauroccvieira/restapi-echo-go/services"

	"github.com/stretchr/testify/assert"
)

type MockUserService struct {
	services.UserService
	MockGetUsers       func() ([]models.User, error)
	MockCreateUser     func(*models.User) (models.User, error)
	MockDeleteUserById func(id int) error
}

func (m *MockUserService) GetUsers() ([]models.User, error) {
	return m.MockGetUsers()
}

func (m *MockUserService) CreateUser(user *models.User) (models.User, error) {
	return m.MockCreateUser(user)
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

func TestCreateUserSuccessCase(t *testing.T) {
	defer setUp_user_test()()

	s := &MockUserService{
		MockCreateUser: func(user *models.User) (models.User, error) {
			return models.User{
				ID:       1,
				Name:     "test",
				Username: "test",
				Password: "test",
			}, nil
		},
	}

	mockService := &services.Services{User: s}

	e := Echo()

	h := New(mockService)

	body := strings.NewReader(`{"name":"test","username":"test","password":"test"}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/user", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.UserHandler.CreateUser(c))
	assert.Equal(t, http.StatusCreated, rec.Code)

}
