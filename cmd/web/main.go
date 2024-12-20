package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Real-Musafir/go-project-booking-reservation/internal/config"
	"github.com/Real-Musafir/go-project-booking-reservation/internal/handlers"
	"github.com/Real-Musafir/go-project-booking-reservation/internal/helpers"
	"github.com/Real-Musafir/go-project-booking-reservation/internal/models"
	"github.com/Real-Musafir/go-project-booking-reservation/internal/render"
	"github.com/alexedwards/scs/v2"
)


const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	


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

func run() error {

	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true ///if close the browser then session should be deleted
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction /// In production it should be true as in production it should be https

	app.Session = session


	tc, err := render.CreateTemplateCache()
	if err!=nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)
	helpers.NewHelplers(&app)

	return nil
}
