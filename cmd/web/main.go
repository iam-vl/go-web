package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/iam-vl/go-web/pkg/config"
	"github.com/iam-vl/go-web/pkg/handlers"
	"github.com/iam-vl/go-web/pkg/render"
)

const portNum = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// Change this to true when if prod
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// True for production
	session.Cookie.Secure = app.InProduction

	app.Session = session

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

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("Starting the app on port %s\n", portNum)

	// If it doesn't work, nothing serious
	// _ = http.ListenAndServe(portNum, nil)
	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
