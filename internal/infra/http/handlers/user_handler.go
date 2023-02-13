package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/test_crud/internal/app"
	"github.com/test_crud/internal/infra/http/requests"
	"github.com/test_crud/internal/infra/http/response"
	"net/http"
)

type RegisterHandler struct {
	as app.AuthService
}

func NewRegisterHandler(a app.AuthService) RegisterHandler {
	return RegisterHandler{
		as: a,
	}
}

func (r RegisterHandler) Register(ctx echo.Context) error {
	var registerUser requests.RegisterAuth
	if err := ctx.Bind(&registerUser); err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Could not decode user data")
	}
	if err := ctx.Validate(&registerUser); err != nil {
		return response.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Could not validate user data")
	}

	userFromRegister := registerUser.RegisterToUser()

	user, err := r.as.Register(userFromRegister)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Could not save new user: %s", err))
	}
	userResponse := user.DomainToResponse()
	return response.Response(ctx, http.StatusCreated, userResponse)
}

func (r RegisterHandler) Login(ctx echo.Context) error {
	var authUser requests.LoginAuth
	if err := ctx.Bind(&authUser); err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Could not decode user data")
	}
	if err := ctx.Validate(&authUser); err != nil {
		return response.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Could not validate user data")
	}
	users, err := r.as.Login(authUser)
	if err != nil {
		return err
	}
	var usersResponse []response.UserResponse
	for _, user := range users {
		resp := user.DomainToResponse()
		usersResponse = append(usersResponse, resp)
	}
	return response.Response(ctx, http.StatusOK, usersResponse)
}
