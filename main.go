package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
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
