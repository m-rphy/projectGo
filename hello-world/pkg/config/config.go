package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// NOTE make sure the you are only importing what is absolutely necessary

// AppConfig holds the application's config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
