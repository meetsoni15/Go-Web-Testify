package main

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
}

func main() {
	// init app config
	app := application{
		Session: getSession(),
	}

	// init router
	mux := app.routes()

	// log message
	log.Println("Server started at port 8080")

	// start server
	if err := http.ListenAndServe(":8080", app.Session.LoadAndSave(mux)); err != nil {
		log.Println(debug.Stack())
		return
	}
}
