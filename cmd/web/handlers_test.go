package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Handler tests
func Test_application_handlers(t *testing.T) {
	theTests := []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"HOME", "/", http.StatusOK},
		{"NOT FOUND", "/meet", http.StatusNotFound},
	}

	// Init application struct
	mux := app.routes() // routes

	// create httptest server
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

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

func Test_application_HandlerHome(t *testing.T) {
	// create req
	req := httptest.NewRequest("GET", "/", nil)

	// create recorder
	rr := httptest.NewRecorder()

	// call context and add it to request
	req = addContextAndSessionToReq(req, app)

	// assign test handler
	testHandler := http.HandlerFunc(app.HandlerHome)

	// call handler
	testHandler.ServeHTTP(rr, req)

	// check status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, but got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "<small>From Session") {
		t.Error("Doesn't contain expected body")
	}
}

func getContext(r *http.Request) context.Context {
	return context.WithValue(r.Context(), contextUserKey, "unknown")
}

func addContextAndSessionToReq(r *http.Request, app application) *http.Request {
	r = r.WithContext(getContext(r))
	ctx, _ := app.Session.Load(r.Context(), r.Header.Get("X-Session"))
	return r.WithContext(ctx)
}
