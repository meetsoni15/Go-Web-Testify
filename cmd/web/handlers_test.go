package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_HandlerHome(t *testing.T) {
	theTests := []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"HOME", "/", http.StatusOK},
		{"NOT FOUND", "/meet", http.StatusNotFound},
	}

	// Init application struct
	var app application
	mux := app.routes() // routes

	// create httptest server
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	pathToTemplate = "../../templates"

	// range through tests
	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("For %s, expected status %d, but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}
