package utils

import (
	"log"
	"renderService/usecase"

	"github.com/gin-gonic/gin"
)

func GetGRPCClient(c *gin.Context) usecase.GRPCClient {
	value, exist := c.Get("GRPCClient")
	if !exist {
		log.Fatalln("no eventService in gin context")
	}
	return value.(usecase.GRPCClient)
}
