package events

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func Routes(db *sqlx.DB) http.Handler {
	repo := NewRepo(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	r := chi.NewRouter()

	r.Get("/", h.list)
	r.Get("/{event_id}", h.get)

	return r
}
