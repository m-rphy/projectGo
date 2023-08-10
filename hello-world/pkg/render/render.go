package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RendersTemplate renders templates useing html/template

// the problem with this approach is that we are reading from
// disk twice (layout and page) for every request
func RenderTemplateSimple(w http.ResponseWriter, gtpl string) {

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

// This will be that data structure that will be out cache
// The key is a string and the value is a pointer to the
// data being store a pointer to template.Template

// NOTE - package level variable
var tc = make(map[string]*template.Template)

// A better version that uses a cache
func RenderTemplate(w http.ResponseWriter, t string) {
	var gtpl *template.Template
	var err error

	// Check to see if we already have
	// the template in the cache
	_, inMap := tc[t] // returns key's value and boolean

	if !inMap {
		// need to create template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cached template")
	}

	gtpl = tc[t]

	err = gtpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(t string) error {
	// parse all the required files for an individual template
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.gtpl",
	}

	// parse the template
	gtpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	// add template to cache
	tc[t] = gtpl
	return nil
}
