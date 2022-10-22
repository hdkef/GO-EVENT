package usecase

import (
	context "context"
	"dbservice/layer"

	"gorm.io/gorm"
)

type LikeService struct {
	*layer.UnimplementedLikeLayerServer
	DB *gorm.DB
}

func (e *LikeService) Get(ctx context.Context, pagination *layer.Pagination) (*layer.LikeList, error) {
	return nil, nil
}

func (e *LikeService) Create(ctx context.Context, payload *layer.Like) (*layer.Empty, error) {
	return nil, nil
}

func (e *LikeService) Delete(ctx context.Context, id *layer.IDPayload) (*layer.Empty, error) {
	return nil, nil
}
