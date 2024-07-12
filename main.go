package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/spf13/viper"
	"github.com/the-Jinxist/tukio-api/config"
	"github.com/the-Jinxist/tukio-api/middleware"
	"github.com/the-Jinxist/tukio-api/pkg/auth"
	"github.com/the-Jinxist/tukio-api/pkg/me"
)

func main() {
	setupConfig()

	db := config.GetDB()

	r := chi.NewRouter()
	r.Use(httprate.LimitByIP(20, 1*time.Minute))
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Mount("/auth", auth.Routes(db))

	r.With(middleware.Authenticator).Mount("/me", me.Routes(db))

	port := viper.GetString("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("running on port: %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)

}

func setupConfig() {
	config.LoadConfigs(".")
	config.InitDB()
}
