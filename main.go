package main

import (
	"fmt"
	"log"
	"net/http"
)

// Home page handler
func Home (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page")
}

// Anout page handler
func About (w http.ResponseWriter, r *http.Request){
	sum := addValues(2,2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is tthe about page and 2+2 is %d", sum))
}

func addValues(x, y int) int {
	sum := x+y
	return sum
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	println("Starting server on :8080")

	// ListenAndServe will block, so handle the error if it fails
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}

	println("This will only print if the server stops unexpectedly.")
}
