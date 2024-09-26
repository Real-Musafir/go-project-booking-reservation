package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alshahadath/go-web/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting server on port %s", portNumber))
	// ListenAndServe will block, so handle the error if it fails
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}

	fmt.Println("This will only print if the server stops unexpectedly.")
}
