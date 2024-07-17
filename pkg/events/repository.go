package events

import (
	"context"
	"database/sql"
	"encoding/base64"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	_ repo = EventsRepo{}

	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

type repo interface {
	list(ctx context.Context, params queryParams) ([]EventResponse, responseParams, error)
	listUserEvents(ctx context.Context, uid string, params queryParams) ([]EventResponse, responseParams, error)
	get(ctx context.Context, eid string) (EventResponse, error)
	create(ctx context.Context, userID string, param createEventParams) error
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
func (e EventsRepo) list(ctx context.Context, params queryParams) ([]EventResponse, responseParams, error) {
	q := psql.Select("events.*, (profiles.first_name || profiles.last_name) as poster_name").
		From("events").Join("profiles on profiles.user_id on events.user_id").
		OrderBy("events.created_at desc").Limit(uint64(params.limit))

	id, _ := base64.StdEncoding.DecodeString(params.cursor)
	eventID := string(id)

	if eventID != "" {
		q = q.Where(sq.Gt{"events.id": eventID})
	}

	query, args, err := q.ToSql()
	if err != nil {
		return nil, responseParams{}, err
	}

	var res []EventResponse
	err = e.db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, responseParams{}, err
	}

	if len(res) == 0 {
		return res, responseParams{}, nil
	}

	lastEvent := res[len(res)-1]
	var resParams responseParams
	resParams.nextCursor = base64.StdEncoding.EncodeToString([]byte(lastEvent.ID.String()))

	return res, resParams, nil

}

func (e EventsRepo) listUserEvents(ctx context.Context, uid string, params queryParams) ([]EventResponse, responseParams, error) {
	q := psql.Select("events.*, (profiles.first_name || profiles.last_name) as poster_name").
		From("events").Join("profiles on profiles.user_id on events.user_id").
		Where(sq.Eq{"events.user_id": uid}).
		OrderBy("events.created_at desc").Limit(uint64(params.limit))

	id, _ := base64.StdEncoding.DecodeString(params.cursor)
	eventID := string(id)

	if eventID != "" {
		q = q.Where(sq.Gt{"events.id": eventID})
	}

	query, args, err := q.ToSql()
	if err != nil {
		return nil, responseParams{}, err
	}

	var res []EventResponse
	err = e.db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, responseParams{}, err
	}

	if len(res) == 0 {
		return res, responseParams{}, nil
	}

	lastEvent := res[len(res)-1]
	var resParams responseParams
	resParams.nextCursor = base64.StdEncoding.EncodeToString([]byte(lastEvent.ID.String()))

	return res, resParams, nil

}

// get implements repo.
func (e EventsRepo) get(ctx context.Context, eid string) (EventResponse, error) {
	var event EventResponse
	err := e.db.GetContext(ctx, &event,
		`select events.*, (profiles.first_name || profiles.last_name) as
		 poster_name from events join profiles on profiles.user_id on events.user_id where events.id = $1`, eid)

	return event, err
}

// create implements repo.
func (e EventsRepo) create(ctx context.Context, userID string, param createEventParams) error {
	tx, err := e.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	eventID := uuid.Must(uuid.NewV7())

	event := Event{
		Name:      param.Name,
		Desc:      param.Desc,
		Picture:   param.Picture,
		Location:  param.Location,
		DressCode: param.DressCode,
		EventTime: param.EventTime,
	}

	_, err = tx.Exec(`insert into events (id, user_id, name, "desc", picture, location, dress_code, event_time, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, now(), now())`,
		eventID, userID, event.Name, event.Desc, event.Picture, event.Location, event.DressCode, event.EventTime)
	if err != nil {
		return err
	}

	if param.TicketCategories != nil {
		if err = insertTicketCategory(tx, eventID, param.TicketCategories); err != nil {
			return err
		}
	}

	return tx.Commit()

}

func insertTicketCategory(tx *sql.Tx, eventID uuid.UUID, ticketCategory []*ticketCategory) error {
	for _, cats := range ticketCategory {
		cate := EventTicketCategory{
			ID:         uuid.Must(uuid.NewV7()),
			EventID:    eventID,
			Name:       cats.Name,
			Desc:       cats.Desc,
			Price:      cats.Price,
			SeatNumber: cats.SeatNumber,
		}

		_, err := tx.Exec(`insert into events_ticket_categories (id, name, "desc", price, event_id, seat_number, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, now(), now())`, cate.ID, cate.Name, cate.Desc, cate.Price, cate.EventID, cate.SeatNumber)
		if err != nil {
			return err
		}
	}

	return nil
}
