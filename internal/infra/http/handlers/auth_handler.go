package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/test_crud/config"
	"github.com/test_crud/internal/app"
	"github.com/test_crud/internal/infra/http/requests"
	"github.com/test_crud/internal/infra/http/response"
	"html/template"
	"log"
	"net/http"
)

type AuthHandler struct {
	as app.AuthService
}

func NewAuthHandler(a app.AuthService) AuthHandler {
	return AuthHandler{
		as: a,
	}
}

func (r AuthHandler) Register(ctx echo.Context) error {
	var registerUser requests.RegisterAuth
	registerUser.Name = ctx.FormValue("name")
	registerUser.Email = ctx.FormValue("email")
	registerUser.Password = ctx.FormValue("password")
	if err := ctx.Validate(&registerUser); err != nil {
		log.Printf("%s: %s", response.ErrorValidateUser, err.Error())
		return response.ErrorResponse(ctx, http.StatusUnprocessableEntity, response.ErrorValidateUser)
	}

	userFromRegister := registerUser.RegisterToUser()

	err := r.as.Register(userFromRegister)
	if err != nil {
		log.Printf("%s: %s", response.ErrorSaveUser, err.Error())
		return response.ErrorResponse(ctx, http.StatusInternalServerError, response.ErrorSaveUser)
	}
	t, _ := template.ParseFiles("temp/result.html", "temp/footer.html", "temp/header.html")
	_ = t.ExecuteTemplate(ctx.Response(), "result", nil)
	return response.MessageResponse(ctx, http.StatusCreated, "User successful created")
}

func (r AuthHandler) Login(ctx echo.Context) error {
	var authUser requests.LoginAuth
	authUser.Email = ctx.FormValue("email")
	authUser.Password = ctx.FormValue("password")
	if err := ctx.Validate(&authUser); err != nil {
		log.Printf("%s: %s", response.ErrorValidateUser, err.Error())
		return response.ErrorResponse(ctx, http.StatusUnprocessableEntity, response.ErrorValidateUser)
	}
	err := r.as.Login(authUser)
	if err != nil {
		log.Printf("%s: %s", response.ErrorLoginUser, err.Error())
		return response.ErrorResponse(ctx, http.StatusInternalServerError, response.ErrorLoginUser)
	}
	return ctx.Redirect(http.StatusTemporaryRedirect, config.GetConfiguration().RedirectUrl)
}
