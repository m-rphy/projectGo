package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Basic func for rendering templates
func renderTemplate(w http.ResponseWriter, gtpl string) {

	parsedTemplate, err := template.ParseFiles("./templates/" + gtpl)
	if err != nil {
		fmt.Println("error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

// Uppercase function makes it public!
// Home is th home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.gtpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.gtpl")
}
