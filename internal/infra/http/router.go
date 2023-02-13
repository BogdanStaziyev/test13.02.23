package http

import (
	"github.com/labstack/echo/v4"
	MW "github.com/labstack/echo/v4/middleware"
	"github.com/test_crud/config/container"
	"github.com/test_crud/internal/infra/http/validators"
)

func EchoRouter(e *echo.Echo, cont container.Container) {

	e.Use(MW.Logger())
	e.Validator = validators.NewValidator()

	//u := e.Group("user")

	v1 := e.Group("/")
	v1.GET("", PingHandler)
}
