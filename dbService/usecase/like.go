package usecase

import (
	context "context"
	"dbservice/layer"
	"dbservice/repository"

	"gorm.io/gorm"
)

type LikeService struct {
	*layer.UnimplementedLikeLayerServer
	DB *gorm.DB
}

func (e *LikeService) Get(ctx context.Context, pagination *layer.Pagination) (*layer.LikeList, error) {
	//create Like repo
	repo := repository.Like{
		DB: e.DB,
	}

	return repo.Get(&ctx, pagination)
}

func (e *LikeService) GetByID(ctx context.Context, id *layer.IDPayload) (*layer.Like, error) {
	//create Like repo
	repo := repository.Like{
		DB: e.DB,
	}

	return repo.GetByID(&ctx, id.ID)
}

func (e *LikeService) Create(ctx context.Context, payload *layer.Like) (*layer.Empty, error) {
	// create Like repo
	repo := repository.Like{
		DB:   e.DB,
		Like: payload,
	}
	//create
	err := repo.Create(&ctx)
	if err != nil {
		return &layer.Empty{}, err
	}
	return &layer.Empty{}, nil
}

func (e *LikeService) Delete(ctx context.Context, id *layer.IDPayload) (*layer.Empty, error) {
	// create Like repo
	repo := repository.Like{
		DB: e.DB,
	}
	//delete
	err := repo.Delete(&ctx, id.ID)
	if err != nil {
		return &layer.Empty{}, err
	}
	return &layer.Empty{}, nil
}
