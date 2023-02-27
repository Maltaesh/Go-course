package main

import (
	"net/http"
	"web3/pkg/config"
	"web3/pkg/handlers"
	"web3/pkg/helpers"
)

func main() {
	var app config.AppConfig

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	helpers.ErrorCheck(err)
}
