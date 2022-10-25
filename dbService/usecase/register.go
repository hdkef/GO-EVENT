package usecase

import (
	context "context"
	"dbservice/layer"
	"dbservice/repository"

	"gorm.io/gorm"
)

type RegisterService struct {
	*layer.UnimplementedRegisterLayerServer
	DB *gorm.DB
}

func (e *RegisterService) Create(ctx context.Context, payload *layer.Register) (*layer.Empty, error) {
	// create register repo
	repo := repository.Register{
		DB:       e.DB,
		Register: payload,
	}
	//create
	err := repo.Create(&ctx)
	if err != nil {
		return &layer.Empty{}, err
	}
	return &layer.Empty{}, nil
}

func (e *RegisterService) GetByID(ctx context.Context, id *layer.IDPayload) (*layer.Register, error) {
	//create register repo
	repo := repository.Register{
		DB: e.DB,
	}

	return repo.GetByID(&ctx, id.ID)
}

func (e *RegisterService) Get(ctx context.Context, pagination *layer.Pagination) (*layer.RegisterList, error) {
	//create register repo
	repo := repository.Register{
		DB: e.DB,
	}

	return repo.Get(&ctx, pagination)
}

func (e *RegisterService) Delete(ctx context.Context, id *layer.IDPayload) (*layer.Empty, error) {
	return nil, nil
}
