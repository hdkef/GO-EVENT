package delivery

import (
	"context"
	"net/http"
	"renderService/layer"
	"renderService/utils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//decode payload
	var user layer.User

	err := utils.DecodeUser(c, &user)
	if err != nil {
		//render error page

		return
	}

	//validate user
	err = utils.ValidateUser(&user)
	if err != nil {
		//render error page

		return
	}

	//hashingPassword
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		//render error page

		return
	}

	//set pass to hashedPassword
	user.Password = string(hashedPass)

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
