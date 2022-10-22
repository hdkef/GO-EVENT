package usecase

import (
	context "context"
	"dbservice/layer"

	"gorm.io/gorm"
)

type RegisterService struct {
	*layer.UnimplementedRegisterLayerServer
	DB *gorm.DB
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
