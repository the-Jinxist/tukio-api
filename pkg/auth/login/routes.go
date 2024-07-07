package login

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
	r.Post("/", h.login)

	return r
}
