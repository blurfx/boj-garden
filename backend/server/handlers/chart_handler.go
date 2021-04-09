package handlers

import (
	"boj-garden/models"
	"boj-garden/server"
	"fmt"
	heatmap "github.com/blurfx/calendar-heatmap"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"time"
)

const days = 7
const weeks = 53

type ChartHandler struct {
	server *server.Server
}

func MakeChartHandler(server *server.Server) *ChartHandler {
	return &ChartHandler{server: server}
}

func (chartHandler *ChartHandler) GetChart(c echo.Context) error {
	user := models.User{}
	submissions := &[]models.Submission{}

	username := c.Param("username")
	today := time.Now()
	firstDay := today.AddDate(0, 0, -(weeks * days) + 1)

	if err := chartHandler.server.DB.First(&user, &models.User{Username: username}).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if err := chartHandler.server.DB.Model(&user).Where("user_id = ? AND date BETWEEN ? AND ?", user.ID, firstDay, today).Association("Submission").Find(&submissions); err != nil {
		fmt.Println(err)
		fmt.Println(submissions)
		return c.JSON(http.StatusNotFound, err)
	}

	data := make(map[heatmap.Date]int)

	for _, submission := range *submissions {
		date := heatmap.Date{
			Year: submission.Date.Year(),
			Month: submission.Date.Month(),
			Day: submission.Date.Day(),
		}
		if _, ok := data[date]; !ok {
			data[date] = 0
		}
		data[date] = int(math.Min(float64(data[date]+1), 4))
	}

	h := heatmap.New(nil)

	buffer := h.Generate(
		heatmap.Date{Year: firstDay.Year(), Month: firstDay.Month(), Day: firstDay.Day()},
		heatmap.Date{Year: today.Year(), Month: today.Month(), Day: today.Day()},
		data,
	)

	return c.Blob(http.StatusOK, "image/svg+xml", buffer.Bytes())
}
