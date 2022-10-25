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

func (u *EventService) Edit(c *gin.Context) error {
	//get param id
	id, err := utils.GetParamID(c)
	if err != nil {
		//send error

		return err
	}

	//decode payload
	var event layer.Event

	selectQ, err := utils.DecodeEvent(c, &event)
	if err != nil {
		//send error

		return err
	}

	//validate event
	err = utils.ValidateEvent(&event, utils.VALIDATE_TYPE_UPDATE)
	if err != nil {
		//send error

		return err
	}

	//get grpc client
	grpc := GetGRPCClient(c)

	//edit
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = grpc.Event.Edit(ctx, &layer.EventEditPayload{ID: id, Select: selectQ, Event: &event})
	if err != nil {
		//send error

		return err
	}
	return nil
}

func (u *EventService) GetByID(c *gin.Context) (*layer.Event, error) {
	//get param id
	id, err := utils.GetParamID(c)
	if err != nil {
		//send error

		return nil, err
	}

	//get grpc client
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpc := GetGRPCClient(c)

	//get event
	return grpc.Event.GetByID(ctx, &layer.IDPayload{ID: id})
}
