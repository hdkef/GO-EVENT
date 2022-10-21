package delivery

import (
	"net/http"
	"renderService/usecase"
	"renderService/utils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func GetAllEvent(c *gin.Context) {
	eventService := utils.GetEventService(c)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	data, err := eventService.GetAllEvent(ctx, &usecase.Pagination{
		LastID: 1,
		Limit:  5,
		Query:  "",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}
