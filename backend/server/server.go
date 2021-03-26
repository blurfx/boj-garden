package server

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
}

func Init() *Server {
	return &Server {
		Echo: echo.New(),
	}
}

func (server *Server) Run(address string) error {
	return server.Echo.Start(address)
}