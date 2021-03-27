package server

import (
	"boj-garden/config"
	"boj-garden/db"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	Echo	*echo.Echo
	DB		*gorm.DB
}

func Init(cfg *config.Config) *Server {
	return &Server {
		Echo:	echo.New(),
		DB:		db.Init(cfg),
	}
}

func (server *Server) Run(address string) error {
	return server.Echo.Start(address)
}