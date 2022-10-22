package usecase

import (
	context "context"
	"dbservice/layer"

	"gorm.io/gorm"
)

type ParticipantService struct {
	*layer.UnimplementedParticipantLayerServer
	DB *gorm.DB
}

func (e *ParticipantService) Get(ctx context.Context, pagination *layer.Pagination) (*layer.ParticipantList, error) {
	return nil, nil
}

func (e *ParticipantService) TagByID(ctx context.Context, payload *layer.TagParticipantPayload) (*layer.Empty, error) {
	return nil, nil
}

func (e *ParticipantService) Delete(ctx context.Context, id *layer.IDPayload) (*layer.Empty, error) {
	return nil, nil
}
