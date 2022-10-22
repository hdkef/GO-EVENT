package delivery

import (
	"renderService/usecase"

	"github.com/gin-gonic/gin"
)

func GRPCMiddleware(m usecase.GRPCClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("GRPCClient", m)
	}
}
