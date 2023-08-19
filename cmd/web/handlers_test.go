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

func Test_application_HandlerHomeTDD(t *testing.T) {
	// test
	theTests := []struct {
		name               string
		sessionValue       string
		expectedResult     string
		expectedStatusCode int
	}{
		{"First Visit", "", "<small>From Session", http.StatusOK},
		{"Second Visit", "unknown", "<small>From Session: unknown", http.StatusOK},
	}

	// iterate through struct
	for _, test := range theTests {
		// create req
		req := httptest.NewRequest("GET", "/", nil)
		// call context and add it to request
		req = addContextAndSessionToReq(req, app)
		// check if session value exists
		// Insert into Session
		if len(test.sessionValue) > 0 {
			// destroy session
			_ = app.Session.Destroy(req.Context())
			// put value inside session
			app.Session.Put(req.Context(), "test", test.sessionValue)
		}

		// create response recorder
		rr := httptest.NewRecorder()

		// create test handler
		testHandler := http.HandlerFunc(app.HandlerHome)

		// call handler
		testHandler.ServeHTTP(rr, req)

		// check status code
		if rr.Code != http.StatusOK {
			t.Errorf("For %s, expected status %d, but got %d", test.name, test.expectedStatusCode, rr.Code)
		}

		// check body
		if !strings.Contains(rr.Body.String(), test.expectedResult) {
			t.Errorf("For %s, expected body contain %s, but it Doesn't contain expected body", test.name, test.expectedResult)
		}

	}
}

// getContext - set and get context
func getContext(r *http.Request) context.Context {
	return context.WithValue(r.Context(), contextUserKey, "unknown")
}

// addContextAndSessionToReq - add context and session to http call
func addContextAndSessionToReq(r *http.Request, app application) *http.Request {
	r = r.WithContext(getContext(r))
	ctx, _ := app.Session.Load(r.Context(), r.Header.Get("X-Session"))
	return r.WithContext(ctx)
}

// render with bad template
func Test_application_render(t *testing.T) {
	theTests := []struct {
		name         string
		templatePath string
		expectedErr  bool
	}{
		{"Template not found", "home.page.html", true},
		{"Execute error", "testdata/test.page.gohtml", true},
	}

	for _, test := range theTests {
		// create response recorder
		rr := httptest.NewRecorder()

		err := app.Render(rr, test.templatePath, &TemplateData{})
		if test.expectedErr && err == nil {
			t.Errorf("For %s expected error but didn't get one", test.name)
		}
	}
}
