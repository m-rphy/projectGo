package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const PORT = ":8080"

// var templatesDir string // Variable to store the absolute path of the templates directory

// func init() {
// 	// Get the absolute path of the directory containing main.go
// 	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
// 	if err != nil {
// 		fmt.Println("error getting current directory:", err)
// 		return
// 	}

// 	// Construct the absolute path to the templates directory
// 	templatesDir = filepath.Join(dir, "templates")
// }

// Uppercase function makes it public!
// Home is th home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.gtpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.gtpl")
}

func renderTemplate(w http.ResponseWriter, gtpl string) {

	parsedTemplate, err := template.ParseFiles("./templates/" + gtpl)
	// parsedTemplate, err := template.ParseFiles(filepath.Join(templatesDir, gtpl))
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

// The basis for any web application in Go!
func main() {

	// Handler functions for different endpoints
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Beep Boop, starting on port%s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
