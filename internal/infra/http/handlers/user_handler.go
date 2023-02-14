package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/test_crud/internal/app"
	"github.com/test_crud/internal/infra/http/response"
	"html/template"
	"net/http"
)

type UserHandler struct {
	us app.UserService
}

func NewUserHandler(a app.UserService) UserHandler {
	return UserHandler{
		us: a,
	}
}

func (r UserHandler) GetUsers(ctx echo.Context) error {
	users, err := r.us.FindAll()
	if err != nil {
		return err
	}
	var usersResponse []response.UserResponse
	for _, user := range users {
		resp := user.DomainToResponse()
		usersResponse = append(usersResponse, resp)
	}
	t, _ := template.ParseFiles("temp/users.html", "temp/footer.html", "temp/header.html")
	_ = t.ExecuteTemplate(ctx.Response(), "users", nil)
	return response.Response(ctx, http.StatusOK, usersResponse)
}
