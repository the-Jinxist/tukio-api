package events

import (
	"time"

	"github.com/gofrs/uuid"
)

type Event struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	Desc      string    `json:"desc" db:"desc"`
	Picture   string    `json:"picture" db:"picture"`
	Location  string    `json:"location" db:"location"`
	DressCode string    `json:"dress_code" db:"dress_code"`
	EventTime time.Time `json:"event_time" db:"event_time"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type EventTicketCategory struct {
	ID         uuid.UUID `json:"id" db:"id"`
	EventID    uuid.UUID `json:"event_id" db:"event_id"`
	Name       string    `json:"name" db:"name"`
	Desc       string    `json:"desc" db:"desc"`
	Price      float64   `json:"price" db:"price"`
	SeatNumber int64     `json:"seat_number" db:"seat_number"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type EventResponse struct {
	ID         uuid.UUID `json:"id" db:"id"`
	UserID     uuid.UUID `json:"user_id" db:"user_id"`
	Name       string    `json:"name" db:"name"`
	PosterName string    `json:"poster_name" db:"poster_name"`
	Desc       string    `json:"desc" db:"desc"`
	Picture    string    `json:"picture" db:"picture"`
	Location   string    `json:"location" db:"location"`
	DressCode  string    `json:"dress_code" db:"dress_code"`
	EventTime  time.Time `json:"event_time" db:"event_time"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type responseParams struct {
	nextCursor string
}
