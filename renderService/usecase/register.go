package usecase

import (
	"context"
	"fmt"
	"renderService/layer"
	"renderService/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterService struct{}

func (u *RegisterService) Create(c *gin.Context, userID *uint32) error {
	//get event id
	eventID, err := utils.GetParamEventID(c)
	if err != nil {
		//send error

		return err
	}

	//decode payload
	var register layer.Register

	_, err = utils.DecodeRegister(c, &register)
	if err != nil {
		//send error

		return err
	}

	//set userID and eventID
	register.Event_ID = eventID
	register.User_ID = userID

	//validate register
	err = utils.ValidateRegister(&register, utils.VALIDATE_TYPE_CREATE)
	if err != nil {
		//send error

		return err
	}

	//get grpc client
	grpc := GetGRPCClient(c)

	//create
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = grpc.Register.Create(ctx, &register)
	if err != nil {
		//send error

		return err
	}

	return nil
}

func (u *RegisterService) GetByID(c *gin.Context) (*layer.Register, error) {
	//get register id
	regID, err := utils.GetParamRegisterID(c)
	if err != nil {
		//send error
		return nil, err
	}

	//get grpc client
	grpc := GetGRPCClient(c)

	//get by id
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return grpc.Register.GetByID(ctx, &layer.IDPayload{ID: regID})
}

func (u *RegisterService) Get(c *gin.Context) (*layer.RegisterList, error) {
	//get event id
	eventID, err := utils.GetParamEventID(c)
	if err != nil {
		//send error
		return nil, err
	}

	//get query param last-id
	lastID, limit, _, err := utils.GetPagination(c)
	if err != nil {
		//send error

		return nil, err
	}

	//get grpc client
	grpc := GetGRPCClient(c)

	//get
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("event_id = %d", *eventID)

	return grpc.Register.Get(ctx, &layer.Pagination{
		LastID: lastID,
		Limit:  limit,
		Query:  &query,
	})
}
