package routes

import (
	"boj-garden/server"
	"boj-garden/server/handlers"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(server *server.Server) {
	registerHandler := handlers.MakeRegisterHandler(server)
	chartHandler := handlers.MakeChartHandler(server)

	server.Echo.Use(middleware.Logger())

	server.Echo.POST("/register", registerHandler.Register)
	server.Echo.GET("/garden/:username", chartHandler.GetChart)
}