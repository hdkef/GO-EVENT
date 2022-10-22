package usecase

import (
	context "context"
	"dbservice/layer"
)

type PaymentService struct {
	*layer.UnimplementedPaymentLayerServer
}

func (e *PaymentService) GetByUserID(ctx context.Context, id *layer.IDPayload) (*layer.PaymentList, error) {
	return nil, nil
}

func (e *PaymentService) GetByID(ctx context.Context, id *layer.IDPayload) (*layer.Payment, error) {
	return nil, nil
}

func (e *PaymentService) Delete(ctx context.Context, id *layer.IDPayload) (*layer.Empty, error) {
	return nil, nil
}

func (e *PaymentService) Create(ctx context.Context, payload *layer.Payment) (*layer.Empty, error) {
	return nil, nil
}
