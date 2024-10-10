package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Real-Musafir/go-project-booking-reservation/internal/config"
	"github.com/Real-Musafir/go-project-booking-reservation/internal/handlers"
	"github.com/Real-Musafir/go-project-booking-reservation/internal/render"
	"github.com/alexedwards/scs/v2"
)


const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true ///if close the browser then session should be deleted
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction /// In production it should be true as in production it should be https

	app.Session = session


	tc, err := render.CreateTemplateCache()
	if err!=nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)


	fmt.Println(fmt.Sprintf("Starting server on port %s", portNumber))
	
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}

	fmt.Println("This will only print if the server stops unexpectedly.")
}
