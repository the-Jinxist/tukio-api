package tickets

import (
	"time"

	"github.com/gofrs/uuid"
)

type Tickets struct {
	ID         uuid.UUID `json:"id" db:"id"`
	UserID     uuid.UUID `json:"user_id" db:"user_id"`
	EventID    uuid.UUID `json:"event_id" db:"event_id"`
	CategoryID uuid.UUID `json:"category_id" db:"category_id"`
	Name       string    `json:"name" db:"name"`
	Desc       string    `json:"desc" db:"desc"`
	Price      float64   `json:"price" db:"price"`
	Quantity   int       `json:"quantity" db:"quantity"`
	Valid      bool      `json:"valid" db:"valid"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type UserTicket struct {
	ID         uuid.UUID `json:"id" db:"id"`
	UserID     uuid.UUID `json:"user_id" db:"user_id"`
	EventID    uuid.UUID `json:"event_id" db:"event_id"`
	EventName  string    `json:"event_name" db:"event_name"`
	CategoryID uuid.UUID `json:"category_id" db:"category_id"`
	Name       string    `json:"name" db:"name"`
	Desc       string    `json:"desc" db:"desc"`
	Price      float64   `json:"price" db:"price"`
	Quantity   int       `json:"quantity" db:"quantity"`
	Valid      bool      `json:"valid" db:"valid"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type UserTicketDetails struct {
	ID         uuid.UUID `json:"id" db:"id"`
	UserID     uuid.UUID `json:"user_id" db:"user_id"`
	EventID    uuid.UUID `json:"event_id" db:"event_id"`
	CategoryID uuid.UUID `json:"category_id" db:"category_id"`
	Name       string    `json:"name" db:"name"`
	Desc       string    `json:"desc" db:"desc"`
	Price      float64   `json:"price" db:"price"`
	Quantity   int       `json:"quantity" db:"quantity"`
	Valid      bool      `json:"valid" db:"valid"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`

	Event
}

type Event struct {
	EventName string    `json:"event_name" db:"event_name"`
	Desc      string    `json:"event_desc" db:"event_desc"`
	Location  string    `json:"event_location" db:"event_location"`
	EventTime time.Time `json:"event_time" db:"event_time"`
}
