package delivery

import (
	"net/http"
	"renderService/layer"
	"renderService/utils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func GetAllEvent(c *gin.Context) {
	grpc := utils.GetGRPCClient(c)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	data, err := grpc.Event.Get(ctx, &layer.Pagination{
		LastID: 1,
		Limit:  5,
		Query:  nil,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}
