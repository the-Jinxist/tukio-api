package events

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/the-Jinxist/tukio-api/middleware"
)

func Routes(db *sqlx.DB) http.Handler {
	repo := NewRepo(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	r := chi.NewRouter()

	r.Get("/public", h.list)
	r.With(middleware.Authenticator).Get("/", h.list)

	r.Get("/{event_id}", h.get)

	r.With(middleware.Authenticator).Get("/your-events", h.listUserEvents)
	r.With(middleware.Authenticator).Post("/", h.create)
	r.With(middleware.Authenticator).Post("/upload", h.uploadImage)

	return r
}
