package handlers

import (
	"net/http"

	"github.com/Real-Musafir/go-project-booking-reservation/pkg/config"
	"github.com/Real-Musafir/go-project-booking-reservation/pkg/models"
	"github.com/Real-Musafir/go-project-booking-reservation/pkg/render"
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

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)


	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// Anout page handler
func (m *Repository) About (w http.ResponseWriter, r *http.Request){

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

