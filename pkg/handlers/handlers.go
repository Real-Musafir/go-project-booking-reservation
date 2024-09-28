package handlers

import (
	"net/http"

	"github.com/alshahadath/go-web/pkg/config"
	"github.com/alshahadath/go-web/pkg/models"
	"github.com/alshahadath/go-web/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository)  Home (w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// Anout page handler
func (m *Repository) About (w http.ResponseWriter, r *http.Request){

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

