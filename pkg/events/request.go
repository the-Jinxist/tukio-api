package events

import "time"

type queryParams struct {
	limit  int
	cursor string
}

type createEventParams struct {
	Name             string            `json:"name" validate:"required,min=10"`
	Desc             string            `json:"desc" validate:"required,min=20"`
	Picture          string            `json:"picture" validate:"required,url"`
	Location         string            `json:"location" validate:"required"`
	DressCode        string            `json:"dress_code"`
	EventTime        time.Time         `json:"event_time" validate:"required"`
	TicketCategories []*ticketCategory `json:"ticket_categories" validate:"dive,required"`
}

type ticketCategory struct {
	Name       string  `json:"ticket_name" validate:"required,min=2"`
	Desc       string  `json:"ticket_desc" validate:"required,min=2"`
	Price      float64 `json:"ticket_price" validate:"required,gt=500"`
	SeatNumber int64   `json:"seat_number" validate:"required,gt=0"`
}
