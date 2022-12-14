package delivery

import (
	"fmt"
	"net/http"
	"renderService/usecase"

	"github.com/gin-gonic/gin"
)

type EventRoute struct{}

func (u *EventRoute) RenderCreate(c *gin.Context) {
	//get user ID //publisher ID from token[TODO]
	id := uint32(1)
	service := usecase.EventService{}
	err := service.Create(c, &id)
	if err != nil {
		//render error
		fmt.Println(err)
		return
	}

	//render ok
	c.JSON(http.StatusOK, "ok")
}

func (u *EventRoute) RenderEdit(c *gin.Context) {
	//get UserID from token [TODO]

	//get event by id [TODO]

	//compare userID and publisherID[TODO]
	service := usecase.EventService{}
	err := service.Edit(c)
	if err != nil {
		//render error

		return
	}

	//render ok
	c.JSON(http.StatusOK, "ok")
}

func (u *EventRoute) RenderGetByID(c *gin.Context) {
	service := usecase.EventService{}
	data, err := service.GetByID(c)
	if err != nil {
		//render error

		return
	}
	//render ok
	c.JSON(http.StatusOK, data)
}

func (u *EventRoute) RenderGet(c *gin.Context) {
	service := usecase.EventService{}
	data, err := service.Get(c)
	if err != nil {
		//render error
		return
	}
	//render ok
	c.JSON(http.StatusOK, data)
}

func (u *EventRoute) AjaxDelete(c *gin.Context) {
	//get user ID from token [TODO]

	//get eventbyID [TODO]

	//compare publisher_id = user ID [TODO]
	service := usecase.EventService{}
	err := service.Delete(c)
	if err != nil {
		//render error
		return
	}
	//render ok
	c.JSON(http.StatusOK, "ok")
}
