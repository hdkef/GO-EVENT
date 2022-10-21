package delivery

import (
	"renderService/usecase"

	"github.com/gin-gonic/gin"
)

func GRPCMiddleware(eventService usecase.EventLayerClient, userService usecase.UserLayerClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("eventService", eventService)
		c.Set("userService", userService)
	}
}
