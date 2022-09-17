package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

// AppConfig holds app configurations
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
