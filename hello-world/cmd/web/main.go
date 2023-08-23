package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/m-rphy/hello-world-go/pkg/config"
	"github.com/m-rphy/hello-world-go/pkg/handlers"
	"github.com/m-rphy/hello-world-go/pkg/render"
)

const PORT = ":8080"

// giving config to middleware
var app config.AppConfig

// variable shadowing
var session *scs.SessionManager

// The basis for any web application in Go!
func main() {

	// change to true when in production
	app.InProduction = false

	// creating a session for clients
	// (by default it uses cookies - but can store in memStore, SQL, etc)
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	// Should the cookies persist after a browswer window closes
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// Usually want it true - allowing https only
	session.Cookie.Secure = app.InProduction
	// Giving session data to other packages
	app.Session = session

	// call Template cache here!
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// setting config file template cache
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// gives render package access to the config
	render.NewTemplates(&app)

	// Start Server
	fmt.Printf("Beep Boop, starting on port%s \n", PORT)

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
