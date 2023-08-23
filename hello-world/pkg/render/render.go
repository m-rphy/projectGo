package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/m-rphy/hello-world-go/pkg/config"
	"github.com/m-rphy/hello-world-go/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, gtpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[gtpl]
	if !ok {
		log.Fatal("could not get template from gtpl cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	// add data to the template
	_ = t.Execute(buf, td)
	// if err != nil {
	// 	log.Println(err)
	// }

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	// get all the files named *.pages.gtpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.gtpl")
	if err != nil {
		return myCache, err
	}

	// range through all the pages
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gtpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gtpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}

////////////////////////////////////////////////////////////////////////////
///////////////////// Less than Ideal methods //////////////////////////////
////////////////////////////////////////////////////////////////////////////

// // This will be that data structure that will be out cache
// // The key is a string and the value is a pointer to the
// // data being store a pointer to template.Template

// // NOTE - package level variable
// var tc = make(map[string]*template.Template)

// // A better version that uses a cache (a simple)
// func RenderTemplateBasic(w http.ResponseWriter, t string) {
// 	var gtpl *template.Template
// 	var err error

// 	// Check to see if we already have
// 	// the template in the cache
// 	_, inMap := tc[t] // returns key's value and boolean

// 	if !inMap {
// 		// need to create template
// 		log.Println("creating template and adding to cache")
// 		err = CreateTemplateCacheBasic(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		log.Println("using cached template")
// 	}

// 	gtpl = tc[t]

// 	err = gtpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func CreateTemplateCacheBasic(t string) error {
// 	// parse all the required files for an individual template
// 	// NOTE - If you have a lot of templates or complex template structure
// 	// this isn't the best. There is another that will do all of this for you
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.gtpl",
// 	}

// 	// parse the template
// 	gtpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	// add template to cache
// 	tc[t] = gtpl
// 	return nil
// }

// // RendersTemplate renders templates useing html/template

// // the problem with this approach is that we are reading from
// // disk twice (layout and page) for every request
// func RenderTemplateWithoutCache(w http.ResponseWriter, gtpl string) {

// 	parsedTemplate, err := template.ParseFiles("./templates/"+gtpl, "./templates/base.layout.gtpl")
// 	if err != nil {
// 		fmt.Println("error parsing template:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	err = parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error parsing template:", err)
// 		return
// 	}
// }
