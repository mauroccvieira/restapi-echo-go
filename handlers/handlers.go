package handlers

import (
	"strings"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mauroccvieira/restapi-echo-go/services"
)

type Handlers struct {
	UserHandler
}

func New(s *services.Services) *Handlers {
	return &Handlers{
		UserHandler: &userHandler{s.User},
	}
}

func SetDefault(e *echo.Echo) {
	e.GET("/healthcheck", HealthCheckHandler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

}

func SetApi(e *echo.Echo, h *Handlers, m echo.MiddlewareFunc) {
	g := e.Group("/api/v1")
	// g.Use(m)

	// Users
	g.GET("/users", h.UserHandler.GetUsers)
	g.POST("/users", h.UserHandler.CreateUser)

}

func Echo() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
		Level: 0,
	}))

	return e
}
