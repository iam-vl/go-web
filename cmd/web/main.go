package main

import (
	"fmt"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
)

const portNum = ":8080"

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache() // Create template cache

	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	render.NewTemplates(&app) // share it with render

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("Starting the app on port %s\n", portNum)

	// If it doesn't work, nothing serious
	_ = http.ListenAndServe(portNum, nil)
}
