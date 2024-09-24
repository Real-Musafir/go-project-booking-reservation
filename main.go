package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

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

func Divide(w http.ResponseWriter, r *http.Request){
	f, err := devideValues(100, 10.0)
	if err!=nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.00, f))
}


func devideValues(x, y float32) (float32, error) {

	if(y<=0){
		err := errors.New("Cannot divide by zero")
		return 0, err
	}

	result := x/y
	return result,nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting server on port %s", portNumber))
	// ListenAndServe will block, so handle the error if it fails
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}

	fmt.Println("This will only print if the server stops unexpectedly.")
}
