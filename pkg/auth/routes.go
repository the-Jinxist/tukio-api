package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/the-Jinxist/tukio-api/pkg/auth/login"
	"github.com/the-Jinxist/tukio-api/pkg/auth/registration"
)

func Routes(db *sqlx.DB) http.Handler {

	r := chi.NewRouter()
	r.Mount("/login", login.Routes(db))
	r.Mount("/register", registration.Routes(db))

	return r
}
