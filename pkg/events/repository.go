package events

import (
	"context"

	"github.com/jmoiron/sqlx"
)

var _ repo = EventsRepo{}

type repo interface {
	list(ctx context.Context, params queryParams) ([]Event, responseParams, error)
	get(ctx context.Context, eid string) (Event, error)
}

type EventsRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) repo {
	return EventsRepo{
		db: db,
	}
}

// list implements repo.
func (e EventsRepo) list(ctx context.Context, params queryParams) ([]Event, responseParams, error) {
	panic("unimplemented")
}

// get implements repo.
func (e EventsRepo) get(ctx context.Context, eid string) (Event, error) {
	panic("unimplemented")
}
