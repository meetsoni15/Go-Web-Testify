package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// init router
	r := chi.NewRouter()

	// Attached middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(app.addIpToContext)

	// register routes
	r.Get("/", app.HandlerHome)
	r.Post("/login", app.HandlerLogin)

	// static assets
	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return r
}
