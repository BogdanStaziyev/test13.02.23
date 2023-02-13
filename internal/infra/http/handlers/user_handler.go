package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/test_crud/internal/app"
	"github.com/test_crud/internal/infra/http/requests"
	"github.com/test_crud/internal/infra/http/response"
	"log"
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
		log.Printf("%s: %s", response.ErrorDecodeUser, err.Error())
		return response.ErrorResponse(ctx, http.StatusBadRequest, response.ErrorDecodeUser)
	}
	if err := ctx.Validate(&registerUser); err != nil {
		log.Printf("%s: %s", response.ErrorValidateUser, err.Error())
		return response.ErrorResponse(ctx, http.StatusUnprocessableEntity, response.ErrorValidateUser)
	}

	userFromRegister := registerUser.RegisterToUser()

	user, err := r.as.Register(userFromRegister)
	if err != nil {
		log.Printf("%s: %s", response.ErrorSaveUser, err.Error())
		return response.ErrorResponse(ctx, http.StatusInternalServerError, response.ErrorSaveUser)
	}
	userResponse := user.DomainToResponse()
	return response.Response(ctx, http.StatusCreated, userResponse)
}

func (r RegisterHandler) Login(ctx echo.Context) error {
	var authUser requests.LoginAuth
	if err := ctx.Bind(&authUser); err != nil {
		log.Printf("%s: %s", response.ErrorDecodeUser, err.Error())
		return response.ErrorResponse(ctx, http.StatusBadRequest, response.ErrorDecodeUser)
	}
	if err := ctx.Validate(&authUser); err != nil {
		log.Printf("%s: %s", response.ErrorValidateUser, err.Error())
		return response.ErrorResponse(ctx, http.StatusUnprocessableEntity, response.ErrorValidateUser)
	}
	users, err := r.as.Login(authUser)
	if err != nil {
		log.Printf("%s: %s", response.ErrorLoginUser, err.Error())
		return response.ErrorResponse(ctx, http.StatusInternalServerError, response.ErrorLoginUser)
	}
	var usersResponse []response.UserResponse
	for _, user := range users {
		resp := user.DomainToResponse()
		usersResponse = append(usersResponse, resp)
	}
	return response.Response(ctx, http.StatusOK, usersResponse)
}
