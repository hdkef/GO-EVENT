package delivery

import (
	"fmt"
	"net/http"
	"renderService/usecase"

	"github.com/gin-gonic/gin"
)

type SubscriptionRoute struct{}

func (u *SubscriptionRoute) AjaxSubs(c *gin.Context) {
	//get user ID from token[TODO]

	id := uint32(3)
	service := usecase.SubscriptionService{}
	err := service.Create(c, &id)
	if err != nil {
		//render error
		fmt.Println(err)
		return
	}

	//render ok
	c.JSON(http.StatusOK, "ok")
}

func (u *SubscriptionRoute) AjaxUnSubs(c *gin.Context) {
	//get consumerID from token[TODO]
	id := uint32(3)

	//get subscription by id
	service := usecase.SubscriptionService{}
	subs, err := service.GetByID(c)
	if err != nil {
		//render error
		fmt.Println(err)
		return
	}

	//compare subscription.consumer_id to consumerID
	if subs.Consumer_ID != nil {
		if *subs.Consumer_ID != id {
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

func (u *SubscriptionRoute) GetByConsumerID(c *gin.Context) {
	//get userID from token[TODO]
	id := uint32(3)

	service := usecase.SubscriptionService{}
	data, err := service.GetByConsumerID(c, &id)
	if err != nil {
		//render error
		fmt.Println(err)
		return
	}

	//render ok
	c.JSON(http.StatusOK, data)
}
