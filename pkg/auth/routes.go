package auth

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-Jinxist/tukio-api/pkg/auth/login"
)

func Routes(db *sql.DB) http.Handler {

	r := chi.NewRouter()
	r.Mount("/login", login.Routes(db))

	return r
}
