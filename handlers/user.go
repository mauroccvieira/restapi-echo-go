package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mauroccvieira/restapi-echo-go/logger"
	"github.com/mauroccvieira/restapi-echo-go/services"
	"github.com/mauroccvieira/restapi-echo-go/utils"
	"go.uber.org/zap"
)

type (
	UserHandler interface {
		GetUsers(c echo.Context) error
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
func (h *userHandler) GetUsers(c echo.Context) error {
	r, err := h.UserService.GetUsers()

	if err != nil {
		logger.Error("failed to get user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}
