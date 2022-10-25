package delivery

import (
	"net/http"
	"renderService/usecase"

	"github.com/gin-gonic/gin"
)

type UserRoute struct{}

func (u *UserRoute) RenderSignIn(c *gin.Context) {
	//if cookies exist and valid, redirect dashboard
	_, err := usecase.ValidateAuthCookie(c)
	if err == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
		return
	}

	service := usecase.UserService{}
	id, err := service.SignIn(c)
	if err != nil {
		//render error
		return
	}
	//create token and set to cookies
	err = usecase.SetAuthCookie(c, id)
	if err != nil {
		//render error
		return
	}

	//render ok
	c.JSON(http.StatusOK, id)
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
	//get userID from token[TODO]
	id := uint32(2)

	service := usecase.UserService{}
	err := service.Edit(c, &id)
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
