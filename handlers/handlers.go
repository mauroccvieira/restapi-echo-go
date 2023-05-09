package handlers

import (
	"strings"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mauroccvieira/restapi-echo-go/services"
)

type UseHandler interface {
	SetApi(e *echo.Echo)
}

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
	setDocs(e)

}

func setDocs(e *echo.Echo) {
	e.GET("/docs", func(c echo.Context) error {
		return c.Redirect(301, "/docs/index.html")
	})
	e.GET("/docs/*", echoSwagger.WrapHandler)
}

func SetApi(e *echo.Echo, h *Handlers, m echo.MiddlewareFunc) {

	hs := make([]UseHandler, 0)

	hs = append(hs, h.UserHandler)

	for _, h := range hs {
		h.SetApi(e)
	}

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
