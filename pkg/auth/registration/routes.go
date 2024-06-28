package registration

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(db *sql.DB) http.Handler {
	repo := NewRepo(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	r := chi.NewRouter()
	r.Post("/", h.register)

	return r
}
