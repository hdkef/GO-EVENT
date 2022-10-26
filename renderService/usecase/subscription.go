package usecase

import (
	"context"
	"fmt"
	"renderService/layer"
	"renderService/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type SubscriptionService struct{}

func (u *SubscriptionService) Create(c *gin.Context, consumerID *uint32) error {
	//get publisher-id
	publisherID, err := utils.GetParamPublisherID(c)
	if err != nil {
		return err
	}

	//get grpc client
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpc := GetGRPCClient(c)

	//create
	_, err = grpc.Subscription.Create(ctx, &layer.Subscription{
		Publisher_ID: publisherID,
		Consumer_ID:  consumerID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *SubscriptionService) Delete(c *gin.Context) error {
	//get subscription-id
	subsID, err := utils.GetParamSubscriptionID(c)
	if err != nil {
		return err
	}

	//get grpc client
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpc := GetGRPCClient(c)

	//delete
	_, err = grpc.Subscription.Delete(ctx, &layer.IDPayload{
		ID: subsID,
	})
	if err != nil {
		return err
	}

	return nil
}

// NEEDED FOR AUTHENTICATING
func (u *SubscriptionService) GetByID(c *gin.Context) (*layer.Subscription, error) {
	//get subscription-id
	subsID, err := utils.GetParamSubscriptionID(c)
	if err != nil {
		return nil, err
	}

	//get grpc client
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpc := GetGRPCClient(c)

	//get by id
	data, err := grpc.Subscription.GetByID(ctx, &layer.IDPayload{
		ID: subsID,
	})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *SubscriptionService) GetByConsumerID(c *gin.Context, consumerID *uint32) (*layer.SubscriptionList, error) {
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
	query := fmt.Sprintf("consumer_id = %d", *consumerID)
	return grpc.Subscription.Get(ctx, &layer.Pagination{
		LastID: lastID,
		Limit:  limit,
		Query:  &query,
	})
}
