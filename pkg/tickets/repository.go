package tickets

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/the-Jinxist/tukio-api/pkg"
)

var (
	_      repo = TicketRepo{}
	sqr         = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	layout      = "2006-01-02 15:04:05.999999 -0700 MST"
)

// Start work on ticket purchase and creation. Think there's a paystack go client
type repo interface {
	listTickets(ctx context.Context, userID string, params pkg.QueryParams) ([]UserTicket, *pkg.ResponseParams, error)
	getTicket(ctx context.Context, userID string, tid string) (UserTicketDetails, error)
}

type TicketRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) repo {
	return TicketRepo{
		db: db,
	}
}

func (t TicketRepo) getTicket(ctx context.Context, userID string, tid string) (UserTicketDetails, error) {

	var res UserTicketDetails
	if err := t.db.GetContext(ctx, &res, `select user_tickets.*, events.name as event_name, events.desc as event_desc, 
		events.location as event_location, event.event_time from user_tickets
		join events on events.id = user_tickets.event_id
		where user_tickets.user_id = $1 and user_tickets.id = $2`, userID); err != nil {
		return res, nil
	}

	return res, nil
}

func (t TicketRepo) listTickets(ctx context.Context, userID string, params pkg.QueryParams) ([]UserTicket, *pkg.ResponseParams, error) {

	q := sqr.Select("user_tickets.*", "events.name as event_name").From("user_tickets").Where(squirrel.Eq{"user_tickets.user_id": userID}).Limit(uint64(params.Limit))
	if params.Cursor != "" {
		id, _ := base64.StdEncoding.DecodeString(params.Cursor)
		eventTime := string(id)

		eTime, _ := time.Parse(layout, eventTime)
		q = q.Where(squirrel.Lt{"user_tickets.created_at": eTime})
	}

	q = q.Join("events on events.id = user_tickets.event_id")
	query, args, err := q.ToSql()
	if err != nil {
		return nil, nil, err
	}

	var tickets []UserTicket
	if err = t.db.SelectContext(ctx, &tickets, query, args...); err != nil {
		return nil, nil, err
	}

	lastTicket := tickets[len(tickets)-1]
	resParams := &pkg.ResponseParams{}
	resParams.NextCursor = base64.StdEncoding.EncodeToString([]byte(lastTicket.CreatedAt.String()))

	return tickets, resParams, nil

}
