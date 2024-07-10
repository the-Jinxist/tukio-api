package events

import (
	"context"
	"encoding/base64"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var (
	_ repo = EventsRepo{}

	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

type repo interface {
	list(ctx context.Context, params queryParams) ([]EventResponse, responseParams, error)
	get(ctx context.Context, eid string) (EventResponse, error)
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

// get implements repo.
func (e EventsRepo) get(ctx context.Context, eid string) (EventResponse, error) {
	var event EventResponse
	err := e.db.GetContext(ctx, &event,
		`select events.*, (profiles.first_name || profiles.last_name) as
		 poster_name from events join profiles on profiles.user_id on events.user_id where events.id = $1`, eid)

	return event, err
}
