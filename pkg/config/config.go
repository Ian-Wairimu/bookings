package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

// AppConfig hold the application configs
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	//InfoLog       *log.Logger
	InProduction bool
	Session      *scs.SessionManager
}
