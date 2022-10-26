package usecase

import (
	context "context"
	"dbservice/layer"
	"dbservice/repository"

	"gorm.io/gorm"
)

type EventService struct {
	*layer.UnimplementedEventLayerServer
	DB *gorm.DB
}

func (e *EventService) Get(ctx context.Context, pagination *layer.Pagination) (*layer.EventList, error) {
	//create event repo
	repo := repository.Event{
		DB: e.DB,
	}

	return repo.Get(&ctx, pagination)
}

func (e *EventService) GetByID(ctx context.Context, id *layer.IDPayload) (*layer.Event, error) {
	//create event repo
	repo := repository.Event{
		DB: e.DB,
	}

	return repo.GetByID(&ctx, id.ID)
}

func (e *EventService) Create(ctx context.Context, payload *layer.Event) (*layer.Empty, error) {
	// create event repo
	repo := repository.Event{
		DB:    e.DB,
		Event: payload,
	}
	//create
	err := repo.Create(&ctx)
	if err != nil {
		return &layer.Empty{}, err
	}
	return &layer.Empty{}, nil
}

func (e *EventService) Delete(ctx context.Context, id *layer.IDPayload) (*layer.Empty, error) {
	//create event repo
	repo := repository.Event{
		DB: e.DB,
	}

	//delete
	err := repo.Delete(&ctx, id.ID)
	if err != nil {
		return &layer.Empty{}, err
	}
	return &layer.Empty{}, nil
}

func (e *EventService) Edit(ctx context.Context, payload *layer.EventEditPayload) (*layer.Empty, error) {
	//create event repo
	repo := repository.Event{
		DB:    e.DB,
		Event: payload.Event,
	}

	//edit
	err := repo.Edit(&ctx, payload.Select, payload.ID)
	if err != nil {
		return &layer.Empty{}, err
	}
	return &layer.Empty{}, nil
}
