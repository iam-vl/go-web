package config

import (
	"log"
	"text/template"
)

// shouldn't import anything strange
type AppConfig struct {
	// UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
