package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplateTEST(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if already have in in the cache
	_, inMap := tc[t]
	if !inMap {
		// need to create
		log.Println("Reading template from disk and adding it to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}

	} else {
		// we have t already
		log.Println("Using cached template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(t string) error {
	line()
	fmt.Println("Starting create template cache")
	fmt.Println("Argument:", t)
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	fmt.Println("Templates: ", templates)

	// Parse template
	fmt.Println("Parsing template(s)")
	tmpl, err := template.ParseFiles(templates...)
	fmt.Println("TMPL: ", tmpl)
	fmt.Printf("TMPL type: %T\n", tmpl)

	if err != nil {
		fmt.Println("ERROR:", err)
		return err
	}

	// Add template to cache
	fmt.Println("Adding template to cache")

	tc[t] = tmpl
	fmt.Println("Entire cache:", tc)
	return nil

}

func line() {
	fmt.Println("==========")
}
