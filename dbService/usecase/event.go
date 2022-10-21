package usecase

import "context"

type EventService struct {
	*UnimplementedEventLayerServer
}

// GetAllEvent will get all event with pagination (lastID, limit, and query)
func (e *EventService) GetAllEvent(ctx context.Context, pagination *Pagination) (*EventList, error) {
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

func (e *EventService) GetAllParticipant(ctx context.Context, pagination *Pagination) (*ParticipantList, error) {
	return nil, nil
}

func (e *EventService) GetEventByID(ctx context.Context, id *IDPayload) (*Event, error) {
	return nil, nil
}

func (e *EventService) RegisterEvent(ctx context.Context, payload *Register) (*Empty, error) {
	return nil, nil
}

func (e *EventService) GetRegisterByUserID(ctx context.Context, id *IDPayload) (*Register, error) {
	return nil, nil
}

func (e *EventService) CreateSubscription(ctx context.Context, payload *Subscription) (*Empty, error) {
	return nil, nil
}

func (e *EventService) GetAllSubscription(ctx context.Context, pagination *Pagination) (*SubscriptionList, error) {
	return nil, nil
}

func (e *EventService) GetPaymentByUserID(ctx context.Context, id *IDPayload) (*PaymentList, error) {
	return nil, nil
}

func (e *EventService) CreatePayment(ctx context.Context, payment *Payment) (*Empty, error) {
	return nil, nil
}
