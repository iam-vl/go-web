package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/iam-vl/go-web/pkg/config"
	"github.com/iam-vl/go-web/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// Multiplexer
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux

}
