package handlers

import (
	"net/http"

	"github.com/alshahadath/go-web/pkg/render"
)

// Home page handler
func Home (w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.html")
}

// Anout page handler
func About (w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.page.html")
}

