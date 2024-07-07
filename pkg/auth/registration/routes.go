package registration

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
	r.Post("/", h.register)
	r.With(middleware.VerifyCodeAuthenticator).Post("/verify_auth", h.verifyCode)

	return r
}
