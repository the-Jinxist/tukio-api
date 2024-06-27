package login

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(db *sql.DB) http.Handler {

	h := handler{}

	r := chi.NewRouter()
	r.Post("/", h.login)

	return r
}
