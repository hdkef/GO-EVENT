package usecase

import (
	"context"
	"renderService/layer"
	"renderService/utils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type EventService struct{}

func (u *EventService) Create(c *gin.Context) error {
	//decode payload
	var event layer.Event

	_, err := utils.DecodeEvent(c, &event)
	if err != nil {
		//send error

		return err
	}

	//validate event
	err = utils.ValidateEvent(&event, utils.VALIDATE_TYPE_CREATE)
	if err != nil {
		//send error

		return err
	}

	//if password exist
	if event.Password != nil {
		//hashingPassword
		hashedByte, err := bcrypt.GenerateFromPassword([]byte(*event.Password), 5)
		if err != nil {
			//send error

			return err
		}

		hashedString := string(hashedByte)

		//set pass to hashedPassword
		event.Password = &hashedString
	}

	//get grpc client
	grpc := GetGRPCClient(c)

	//sign up
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = grpc.Event.Create(ctx, &event)
	if err != nil {
		//send error

		return err
	}

	return nil
}
