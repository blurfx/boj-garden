package handlers

import (
	"boj-garden/models"
	"boj-garden/requests"
	"boj-garden/server"
	"boj-garden/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RegisterHandler struct {
	server *server.Server
}

func MakeRegisterHandler(server *server.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

func (registerHandler *RegisterHandler) Register(c echo.Context) error {
	registerRequest := new(requests.UserRequest)

	if err := c.Bind(registerRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(registerRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	username := registerRequest.Username
	user := &models.User{}
	registerHandler.server.DB.FirstOrCreate(&user, &models.User{Username: username})

	crawlTask := &utils.CrawlTask{
		DB: registerHandler.server.DB,
		User: user,
	}
	utils.EnqueueCrawlTask(crawlTask)

	return c.JSON(http.StatusOK, registerRequest)
}