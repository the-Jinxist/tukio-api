package auth

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-Jinxist/tukio-api/pkg/auth/login"
	"github.com/the-Jinxist/tukio-api/pkg/auth/registration"
)

func Routes(db *sql.DB) http.Handler {

	r := chi.NewRouter()
	r.Mount("/login", login.Routes(db))
	r.Mount("/register", registration.Routes(db))

	return r
}
