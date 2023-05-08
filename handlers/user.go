package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mauroccvieira/restapi-echo-go/handlers/requests"
	"github.com/mauroccvieira/restapi-echo-go/logger"
	"github.com/mauroccvieira/restapi-echo-go/models"
	"github.com/mauroccvieira/restapi-echo-go/services"
	"github.com/mauroccvieira/restapi-echo-go/utils"
	"go.uber.org/zap"
)

type (
	UserHandler interface {
		GetUsers(c echo.Context) error
		CreateUser(c echo.Context) error
	}

	userHandler struct {
		services.UserService
	}
)

// GetUsers
//
//	@Summary		Fetch a list of all users.
//	@Description	Fetch a list of all users.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	[]models.User
//	@Failure		500	{object}	utils.Error
//	@Router			/api/v1/user [get]
//
// @paramete
func (h *userHandler) GetUsers(c echo.Context) error {
	r, err := h.UserService.GetUsers()

	if err != nil {
		logger.Error("failed to get user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

// CreateUser
// @Summary Create an user.
// @Description Create an user.
// @Tags User
// @Accept json
// @Produce json
// @Param user body requests.CreateUserRequest true "User"
// @Success 201 {object} responses.CreateUserResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /api/v1/user [post]
func (h *userHandler) CreateUser(c echo.Context) error {
	var req *requests.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		logger.Error("failed to bind user", zap.Error(err))
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}

	user := &models.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}

	r, err := h.UserService.CreateUser(user)

	if err != nil {
		logger.Error("failed to create user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, r)
}
