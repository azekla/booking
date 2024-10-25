package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/azekla/booking/pkg/config"
	"github.com/azekla/booking/pkg/handlers"
	"github.com/azekla/booking/pkg/render"
	"log"
	"net/http"
	"time"
)

// set the port number for the server
const portNumber = ":7182"

var app config.AppConfig
var session *scs.SessionManager

// the main function
func main() {
	// change this to true when in production
	app.InProduction = false
 	// create a new session manager
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	// assign the session manager to the app config
	app.Session = session

	// create a new template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// assign the template cache to the app config
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
