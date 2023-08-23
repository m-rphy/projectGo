package handlers

import (
	"net/http"

	"github.com/m-rphy/hello-world-go/pkg/config"
	"github.com/m-rphy/hello-world-go/pkg/models"
	"github.com/m-rphy/hello-world-go/pkg/render"
)

// NOTE Uppercase function names are public!

// Repo is the repository use by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// grab the remote IP and store it in the session
	// Returned as a string
	remoteIP := r.RemoteAddr
	// put the ip in the session
	m.App.Session.Put(r.Context(), "Name_of_IP", remoteIP)

	render.RenderTemplate(w, "home.page.gtpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// Test rendering Data on the about page dynamically
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	// Get IP address from session and put in StringMap to render on the page
	remoteIP := m.App.Session.GetString(r.Context(), "Name_of_IP")
	stringMap["remoteIP"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, "about.page.gtpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
