package usecase

import (
	context "context"
	"dbservice/layer"

	"gorm.io/gorm"
)

type EventService struct {
	*layer.UnimplementedEventLayerServer
	DB *gorm.DB
}

func (e *EventService) Get(ctx context.Context, pagination *layer.Pagination) (*layer.EventList, error) {
	nama := "marathon"
	desc := "event tahunan marathon di bandung"
	needPayment := false
	needID := false
	isOffline := true
	presenceQ := "{`q`:[]}"
	status := layer.EventStatus_E_OPEN_FOR_REG

	list := []*layer.Event{
		{
			Name:             nama,
			Desc:             desc,
			NeedPayment:      needPayment,
			Need_ID:          needID,
			IsOffline:        isOffline,
			PresenceQuestion: presenceQ,
			Status:           status,
		},
	}

	return &layer.EventList{
		List: list,
	}, nil
}

func (e *EventService) GetByID(ctx context.Context, id *layer.IDPayload) (*layer.Event, error) {
	return nil, nil
}

func (e *EventService) Create(ctx context.Context, payload *layer.Event) (*layer.Empty, error) {
	return nil, nil
}

func (e *EventService) Delete(ctx context.Context, id *layer.IDPayload) (*layer.Empty, error) {
	return nil, nil
}

func (e *EventService) Edit(ctx context.Context, payload *layer.EventEditPayload) (*layer.Empty, error) {
	return nil, nil
}
