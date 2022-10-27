package delivery

import (
	"fmt"
	"net/http"
	"renderService/usecase"

	"github.com/gin-gonic/gin"
)

type CertificateRoute struct{}

func (u *CertificateRoute) AjaxCreate(c *gin.Context) {
	//get UserID from token [TODO]

	//get event by id [TODO]

	//compare event.Publisher = userID from token

	//create
	service := usecase.CertificateService{}
	err := service.Create(c)
	if err != nil {
		//render error
		fmt.Println(err)
		return
	}

	//render ok
	c.JSON(http.StatusOK, "ok")
}

func (u *CertificateRoute) RenderGetByUserID(c *gin.Context) {
	//get user ID from token[TODO]

	id := uint32(3)
	service := usecase.CertificateService{}
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
