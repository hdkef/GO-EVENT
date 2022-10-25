package delivery

import (
	"fmt"
	"net/http"
	"renderService/usecase"

	"github.com/gin-gonic/gin"
)

type RegisterRoute struct{}

func (u *RegisterRoute) RenderGetByID(c *gin.Context) {
	//get user ID from token[TODO]

	//get register
	service := usecase.RegisterService{}
	data, err := service.GetByID(c)
	if err != nil {
		//render error

		return
	}
	//compare register.user_id == user ID [TODO]
	//render ok
	c.JSON(http.StatusOK, data)
}

func (u *RegisterRoute) RenderCreate(c *gin.Context) {
	//get user ID from token[TODO]
	id := uint32(1)
	service := usecase.RegisterService{}
	err := service.Create(c, &id)
	if err != nil {
		//render error
		fmt.Println(err)
		return
	}

	//render ok
	c.JSON(http.StatusOK, "ok")
}

func (u *RegisterRoute) RenderGetByEventID(c *gin.Context) {
	//get user ID from token[TODO]

	//get event by id[TODO]

	//compare if event.creator id = user ID[TODO]

	service := usecase.RegisterService{}
	data, err := service.Get(c)
	if err != nil {
		//render error

		return
	}
	//render ok
	c.JSON(http.StatusOK, data)
}
