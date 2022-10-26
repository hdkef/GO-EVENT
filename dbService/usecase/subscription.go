package usecase

import (
	context "context"
	"dbservice/layer"
	"dbservice/repository"

	"gorm.io/gorm"
)

type SubscriptionService struct {
	*layer.UnimplementedSubscriptionLayerServer
	DB *gorm.DB
}

func (e *SubscriptionService) GetByID(ctx context.Context, id *layer.IDPayload) (*layer.Subscription, error) {
	//create Subscription repo
	repo := repository.Subscription{
		DB: e.DB,
	}

	return repo.GetByID(&ctx, id.ID)
}

func (e *SubscriptionService) Get(ctx context.Context, pagination *layer.Pagination) (*layer.SubscriptionList, error) {
	//create Subscription repo
	repo := repository.Subscription{
		DB: e.DB,
	}

	return repo.Get(&ctx, pagination)
}

func (e *SubscriptionService) Delete(ctx context.Context, id *layer.IDPayload) (*layer.Empty, error) {
	// create Subscription repo
	repo := repository.Subscription{
		DB: e.DB,
	}
	//delete
	err := repo.Delete(&ctx, id.ID)
	if err != nil {
		return &layer.Empty{}, err
	}
	return &layer.Empty{}, nil
}

func (e *SubscriptionService) Create(ctx context.Context, payload *layer.Subscription) (*layer.Empty, error) {
	// create Subscription repo
	repo := repository.Subscription{
		DB:           e.DB,
		Subscription: payload,
	}
	//create
	err := repo.Create(&ctx)
	if err != nil {
		return &layer.Empty{}, err
	}
	return &layer.Empty{}, nil
}
