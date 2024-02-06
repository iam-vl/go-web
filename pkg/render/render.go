package render

import (
	"bytes"
	"log"
	"myapp/pkg/config"
	"net/http"
	"path/filepath"
	"text/template"
)

// var functions = template.FuncMap{}

// handlers and render need the app config from min
var app *config.AppConfig

// Sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Get template cache from the app config
	tc := app.TemplateCache

	// Get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Cannot get template from template cache")
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil) // Ignore error

	// Render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {

		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		}

		myCache[name] = ts
	}
	return myCache, nil
}
