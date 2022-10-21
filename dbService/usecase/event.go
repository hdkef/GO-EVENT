package usecase

import "context"

type EventService struct {
}

// GetAllEvent will get all event with pagination (lastID, limit, and query)
func (e *EventService) GetAllEvent(ctx *context.Context, pagination *Pagination) (*EventList, error) {
	ID := uint32(1)
	Name := "Marathon"
	Desc := "Pocari Sweat Marathon 42 km in Bandung"
	NeedPayment := false
	Need_ID := false
	Payment_ID := uint32(1)
	Creator_ID := uint32(1)
	example := Event{ID: &ID, Name: &Name, Desc: &Desc, NeedPayment: &NeedPayment, Need_ID: &Need_ID, Payment_ID: &Payment_ID, Creator_ID: &Creator_ID}

	return &EventList{
		List: []*Event{&example, &example, &example},
	}, nil
}
