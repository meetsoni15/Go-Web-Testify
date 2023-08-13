package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_addIpToContext(t *testing.T) {
	tests := []struct {
		// header
		headerName  string
		headerValue string
		// addr
		addr string
		// empty address
		emptyAddr bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "10.68.101.31", "", false},
		{"", "", "hello:world", false},
	}

	// init application
	app := application{}

	// create dummy handler to check context values
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get context value
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not found")
		}

		// type cast val to string
		ip, ok := val.(string)
		if !ok {
			t.Error("Not string")
		}

		t.Log(ip)
	})

	// itreate thourgh test cases
	for _, e := range tests {
		testHandler := app.addIpToContext(nextHandler)
		// create dummy request
		req := httptest.NewRequest("GET", "https://meetsoni.me", nil)
		if len(e.headerName) > 0 {
			req.Header.Set(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}

		if e.emptyAddr {
			req.RemoteAddr = ""
		}

		testHandler.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {
	// init application
	app := application{}

	// define empty context
	ctx := context.Background()

	// assign value to context
	ctx = context.WithValue(ctx, contextUserKey, "10.68.101.31")

	// call ip from context func
	if !strings.EqualFold(app.ipFromContext(ctx), "10.68.101.31") {
		t.Error("Wrong value returned from context")
	}
}
