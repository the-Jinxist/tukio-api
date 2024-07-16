package events

import (
	"context"

	"github.com/the-Jinxist/tukio-api/middleware"
)

var _ service = EventsService{}

type service interface {
	rlist(ctx context.Context, params queryParams) ([]EventResponse, responseParams, error)
	get(ctx context.Context, eid string) (EventResponse, error)
	listUserEvents(ctx context.Context, params queryParams) ([]EventResponse, responseParams, error)
	create(ctx context.Context, param createEventParams) error
}

type EventsService struct {
	repo repo
}

func NewService(r repo) service {
	return EventsService{repo: r}
}

// get implements service.
func (e EventsService) get(ctx context.Context, eid string) (EventResponse, error) {
	return e.repo.get(ctx, eid)
}

// rlist implements service.
func (e EventsService) rlist(ctx context.Context, params queryParams) ([]EventResponse, responseParams, error) {
	if middleware.GetUserID(ctx) == "" {
		params.limit = 5
		params.cursor = ""
	}

	return e.repo.list(ctx, params)

}

// listUserEvents implements service.
func (e EventsService) listUserEvents(ctx context.Context, params queryParams) ([]EventResponse, responseParams, error) {
	userID := middleware.GetUserID(ctx)
	return e.repo.listUserEvents(ctx, userID, params)
}

// create implements service.
func (e EventsService) create(ctx context.Context, param createEventParams) error {
	userID := middleware.GetUserID(ctx)
	return e.repo.create(ctx, userID, param)
}
