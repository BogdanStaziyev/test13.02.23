package http

import "github.com/labstack/echo/v4"

type Server struct {
	Echo *echo.Echo
}

func NewServer() *Server {
	s := &Server{
		Echo: echo.New(),
	}
	return s
}

func (s Server) Start(port string) error {
	return s.Echo.Start(":" + port)
}
