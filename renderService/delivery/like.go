package delivery

import (
	"fmt"
	"net/http"
	"renderService/usecase"

	"github.com/gin-gonic/gin"
)

type LikeRoute struct{}

func (u *LikeRoute) RenderGetByEventID(c *gin.Context) {
	service := usecase.LikeService{}
	data, err := service.GetByEventID(c)
	if err != nil {
		//render error
		fmt.Println(err)
		return
	}

	//render ok
	c.JSON(http.StatusOK, data)
}

func (u *LikeRoute) RenderGetByUserID(c *gin.Context) {
	//get user ID from token[TODO]

	id := uint32(3)
	service := usecase.LikeService{}
	//compare like.UserID == UserID [TODO]
	data, err := service.GetByUserID(c, &id)
	if err != nil {
		//render error
		fmt.Println(err)
		return
	}

	//render ok
	c.JSON(http.StatusOK, data)
}

func (u *LikeRoute) AjaxCreate(c *gin.Context) {
	//get user ID from token[TODO]

	id := uint32(3)
	service := usecase.LikeService{}
	err := service.Create(c, &id)
	if err != nil {
		//render error
		fmt.Println(err)
		return
	}

	//render ok
	c.JSON(http.StatusOK, "ok")
}

func (u *LikeRoute) AjaxDelete(c *gin.Context) {
	//get consumerID from token[TODO]
	id := uint32(3)

	//get like by id
	service := usecase.LikeService{}
	like, err := service.GetByID(c)
	if err != nil {
		//render error
		fmt.Println(err)
		return
	}

	//compare Like.User_ID to userID
	if like.User_ID != nil {
		if *like.User_ID != id {
			//render error page
			return
		}
	}

	//delete
	err = service.Delete(c)
	if err != nil {
		//render error
		fmt.Println(err)
		return
	}

	//render ok
	c.JSON(http.StatusOK, "ok")
}
