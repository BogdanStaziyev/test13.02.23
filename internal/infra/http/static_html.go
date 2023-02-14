package http

import (
	"github.com/labstack/echo/v4"
	"html/template"
)

func register(ctx echo.Context) error {
	t, _ := template.ParseFiles("temp/register.html", "temp/footer.html", "temp/header.html")
	err := t.ExecuteTemplate(ctx.Response(), "main", nil)
	if err != nil {
		return err
	}
	return nil
}

func login(ctx echo.Context) error {
	t, _ := template.ParseFiles("temp/login.html", "temp/footer.html", "temp/header.html")
	err := t.ExecuteTemplate(ctx.Response(), "main", nil)
	if err != nil {
		return err
	}
	return nil
}
