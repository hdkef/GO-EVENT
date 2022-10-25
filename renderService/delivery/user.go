package delivery

import (
	"net/http"
	"renderService/usecase"

	"github.com/gin-gonic/gin"
)

type UserRoute struct{}

func (u *UserRoute) RenderSignIn(c *gin.Context) {
	service := usecase.UserService{}
	err := service.SignIn(c)
	if err != nil {
		//render error
		return
	}
	//create token
	s := "initokenpurapurayangdibikinuntuktesting"

	//set token to cookies

	//render ok
	c.JSON(http.StatusOK, s)
}

func (u *UserRoute) RenderSignUp(c *gin.Context) {
	service := usecase.UserService{}
	err := service.SignUp(c)
	if err != nil {
		//render error

		return
	}
	//render ok
	c.JSON(http.StatusOK, "ok")
}

func (u *UserRoute) RenderEdit(c *gin.Context) {
	service := usecase.UserService{}
	err := service.Edit(c)
	if err != nil {
		//render error

		return
	}
	//render ok
	c.JSON(http.StatusOK, "ok")
}

func (u *UserRoute) RenderGetByID(c *gin.Context) {
	service := usecase.UserService{}
	data, err := service.GetByID(c)
	if err != nil {
		//render error

		return
	}
	//render ok
	c.JSON(http.StatusOK, data)
}
