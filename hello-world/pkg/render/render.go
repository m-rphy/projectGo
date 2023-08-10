package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RendersTemplate renders templates useing html/template
func RenderTemplate(w http.ResponseWriter, gtpl string) {

	parsedTemplate, err := template.ParseFiles("./templates/"+gtpl, "./templates/base.layout.gtpl")
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
