package response

import (
	"github.com/labstack/echo/v4"
)

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(c echo.Context, statusCode int, data interface{}) error {
	return c.JSON(statusCode, data)
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Error{
		Code:  statusCode,
		Error: message,
	})
}

const (
	ErrorDecodeUser   = "Could not decode user data"
	ErrorValidateUser = "Could not validate user data"
	ErrorSaveUser     = "Could not save, user already exist"
	ErrorLoginUser    = "Could not login user invalid email or password"
)
