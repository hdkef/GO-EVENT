package usecase

import (
	"context"
	"fmt"
	"renderService/layer"
	"renderService/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type LikeService struct{}

func (u *LikeService) Create(c *gin.Context, userID *uint32) error {
	//get event id
	eventID, err := utils.GetParamEventID(c)
	if err != nil {
		return err
	}

	//get grpc client
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpc := GetGRPCClient(c)

	//create
	_, err = grpc.Like.Create(ctx, &layer.Like{User_ID: userID, Event_ID: eventID})
	if err != nil {
		return err
	}

	return nil
}

func (u *LikeService) Delete(c *gin.Context) error {
	//get Like-id
	likeID, err := utils.GetParamLikeID(c)
	if err != nil {
		return err
	}

	//get grpc client
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpc := GetGRPCClient(c)

	//delete
	_, err = grpc.Like.Delete(ctx, &layer.IDPayload{
		ID: likeID,
	})
	if err != nil {
		return err
	}

	return nil
}

// UNTUK AUTHENTICATION
func (u *LikeService) GetByID(c *gin.Context) (*layer.Like, error) {
	//get like id
	likeID, err := utils.GetParamLikeID(c)
	if err != nil {
		return nil, err
	}

	//get grpc client
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpc := GetGRPCClient(c)

	//get by id
	data, err := grpc.Like.GetByID(ctx, &layer.IDPayload{ID: likeID})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *LikeService) GetByUserID(c *gin.Context, userID *uint32) (*layer.LikeList, error) {
	//get pagination
	lastID, limit, _, err := utils.GetPagination(c)
	if err != nil {
		return nil, err
	}

	//get grpc client
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpc := GetGRPCClient(c)

	//get
	query := fmt.Sprintf("user_id = %d", *userID)
	return grpc.Like.Get(ctx, &layer.Pagination{
		LastID: lastID,
		Limit:  limit,
		Query:  &query,
	})
}

func (u *LikeService) GetByEventID(c *gin.Context) (*layer.LikeList, error) {
	//get event id
	eventID, err := utils.GetParamEventID(c)
	if err != nil {
		return nil, err
	}

	//get pagination
	lastID, limit, _, err := utils.GetPagination(c)
	if err != nil {
		return nil, err
	}

	//get grpc client
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpc := GetGRPCClient(c)

	//get
	query := fmt.Sprintf("event_id = %d", *eventID)
	return grpc.Like.Get(ctx, &layer.Pagination{
		LastID: lastID,
		Limit:  limit,
		Query:  &query,
	})
}
