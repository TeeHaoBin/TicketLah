package repositories

import (
	"context"
	"github.com/TeeHaoBin/TicketLah/backend/models"
	"time"
)

type EventRepository struct {
	db any
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	// mock data
	events := []*models.Event{}

	events = append(events, &models.Event{
		ID: "1",
		Name: "Concert A",
		Location: "Stadium X",
		Date:  time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId string) ([]*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) ([]*models.Event, error) {
	return nil, nil
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository {
		db: db,
	}
}