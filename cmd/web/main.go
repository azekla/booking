package main

import (
	"fmt"
	"github.com/azekla/booking/pkg/config"
	"github.com/azekla/booking/pkg/handlers"
	"github.com/azekla/booking/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":7182"

// the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
	}

}
