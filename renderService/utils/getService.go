package utils

import (
	"log"
	"renderService/usecase"

	"github.com/gin-gonic/gin"
)

func GetEventService(c *gin.Context) usecase.EventLayerClient {
	value, exist := c.Get("eventService")
	if !exist {
		log.Fatalln("no eventService in gin context")
	}
	return value.(usecase.EventLayerClient)
}
