package me

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
	r.Get("/", h.get)
	r.Put("/", h.update)
	r.Get("/profile/{user_id}", h.getUserProfile)

	return r
}
