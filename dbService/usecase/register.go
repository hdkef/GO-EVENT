package usecase

import (
	context "context"
	"dbservice/layer"
)

type RegisterService struct {
	*layer.UnimplementedRegisterLayerServer
}

func (e *RegisterService) Create(ctx context.Context, payload *layer.Register) (*layer.Empty, error) {
	return nil, nil
}

func (e *RegisterService) GetByID(ctx context.Context, id *layer.IDPayload) (*layer.Register, error) {
	return nil, nil
}

func (e *RegisterService) Delete(ctx context.Context, id *layer.IDPayload) (*layer.Empty, error) {
	return nil, nil
}
