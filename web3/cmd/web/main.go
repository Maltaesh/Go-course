package main

import (
	"net/http"
	"time"
	"web3/pkg/config"
	"web3/pkg/handlers"
	"web3/pkg/helpers"

	"github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager
var app config.AppConfig

func main() {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	// cookies secure is falce because in development there is no https connection
	sessionManager.Cookie.Secure = false
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = sessionManager

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	helpers.ErrorCheck(err)
}
