package delivery

import (
	"net/http"
	"renderService/usecase"

	"github.com/gin-gonic/gin"
)

type EventRoute struct{}

func (u *EventRoute) Create(c *gin.Context) {
	service := usecase.EventService{}
	err := service.Create(c)
	if err != nil {
		//render error
		return
	}

	//render ok
	c.JSON(http.StatusOK, "ok")
}
