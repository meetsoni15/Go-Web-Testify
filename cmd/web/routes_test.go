package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_application_router(t *testing.T) {
	registered := []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/static/*", "GET"},
	}

	var app application
	mux := app.routes()

	// typecase mux to chi.Routes
	chiRoutes := mux.(chi.Routes)

	for _, r := range registered {
		if !routeExists(r.route, r.method, chiRoutes) {
			t.Errorf("Route %s with method %s is not registered", r.route, r.method)
		}
	}
}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false
	// walk interface
	_ = chi.Walk(chiRoutes, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(testMethod, method) && strings.EqualFold(testRoute, route) {
			found = true
		}
		return nil
	})

	return found
}
