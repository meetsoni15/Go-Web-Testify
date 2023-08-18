package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// init router
	r := chi.NewMux()

	// Attached middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(app.addIpToContext)
	// session manager middleware
	r.Use(app.Session.LoadAndSave)

	// register routes
	r.Get("/", app.HandlerHome)
	r.Post("/login", app.HandlerLogin)

	// static assets
	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return r
}
