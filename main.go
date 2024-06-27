package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"github.com/the-Jinxist/tukio-api/config"
	"github.com/the-Jinxist/tukio-api/pkg/auth"
)

func main() {
	setupConfig()

	db := config.GetDB()

	r := chi.NewRouter()
	r.Mount("/auth", auth.Routes(db))

	port := viper.GetString("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)

}

func setupConfig() {
	config.LoadConfigs(".")
	config.InitDB()
}
