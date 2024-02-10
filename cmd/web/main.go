package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iam-vl/go-web/pkg/config"
	"github.com/iam-vl/go-web/pkg/handlers"
	"github.com/iam-vl/go-web/pkg/render"
)

const portNum = ":8080"

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache() // Create template cache

	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	// Create a repo var and pass it bck to handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app) // share it with render

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("Starting the app on port %s\n", portNum)

	// If it doesn't work, nothing serious
	_ = http.ListenAndServe(portNum, nil)
}
