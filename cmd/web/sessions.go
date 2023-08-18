package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

// getSession will create new session and return it
func getSession() *scs.SessionManager {
	log.Println("Here")
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	return session
}
