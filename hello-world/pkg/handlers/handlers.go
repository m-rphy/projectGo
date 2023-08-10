package handlers

import (
	"net/http"

	"github.com/m-rphy/hello-world-go/pkg/render"
)

// Uppercase function makes it public!
// Home is th home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gtpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gtpl")
}
