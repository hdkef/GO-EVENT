package delivery

import (
	"context"
	"net/http"
	"renderService/layer"
	"renderService/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	//decode payload
	var user layer.User

	err := utils.DecodeUser(c, &user)
	if err != nil {
		//render error page

		return
	}

	//get grpc client
	grpc := utils.GetGRPCClient(c)

	//sign up
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = grpc.User.Create(ctx, &user)
	if err != nil {
		//render error page

		return
	}

	//render ok page
	c.JSON(http.StatusOK, "ok")
}
