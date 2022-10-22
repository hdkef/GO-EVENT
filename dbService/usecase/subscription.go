package usecase

import (
	context "context"
	"dbservice/layer"
)

type SubscriptionService struct {
	*layer.UnimplementedSubscriptionLayerServer
}

func (e *SubscriptionService) Get(ctx context.Context, pagination *layer.Pagination) (*layer.SubscriptionList, error) {
	return nil, nil
}

func (e *SubscriptionService) Delete(ctx context.Context, id *layer.IDPayload) (*layer.Empty, error) {
	return nil, nil
}

func (e *SubscriptionService) Create(ctx context.Context, payload *layer.Subscription) (*layer.Empty, error) {
	return nil, nil
}
