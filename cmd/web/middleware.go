package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

// concerete type
type contextKey string

// context key
const contextUserKey contextKey = "user_ip"

// ipFromContext -> Get ip address from request context
func (app *application) ipFromContext(ctx context.Context) string {
	return ctx.Value(contextUserKey).(string)
}

// addIpToContext is middleware
func (app *application) addIpToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// init ctx
		var ctx = context.Background()
		// get most absolute ip
		ip, err := getIP(r)
		if err != nil {
			ip, _, _ = net.SplitHostPort(r.RemoteAddr)
			if len(ip) == 0 {
				ip = "unknown"
			}
		}
		ctx = context.WithValue(r.Context(), contextUserKey, ip)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// getIP used to get real ip
// weather its coming as forward proxy or something else
func getIP(r *http.Request) (string, error) {
	// get ip
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "unknown", err
	}

	// parse ip to check if its valid or not
	userIP := net.ParseIP(ip)
	if userIP == nil {
		return "", fmt.Errorf("userIP: %q is not in IP:PORT format", r.RemoteAddr)
	}

	// check if its forwarded
	forwarded := r.Header.Get("X-Forwarded-For")
	if len(forwarded) > 0 {
		ip = forwarded
	}

	if len(ip) == 0 {
		ip = "forwarded"
	}

	return ip, nil
}
