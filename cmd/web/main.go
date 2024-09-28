package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alshahadath/go-web/pkg/config"
	"github.com/alshahadath/go-web/pkg/handlers"
	"github.com/alshahadath/go-web/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err!=nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting server on port %s", portNumber))
	// ListenAndServe will block, so handle the error if it fails
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}

	fmt.Println("This will only print if the server stops unexpectedly.")
}
