package main

import (
	"log"
	"net/http"
	"runtime/debug"
)

type application struct{}

func main() {
	// init app config
	app := application{}

	// init router
	mux := app.routes()

	// log message
	log.Println("Server started at port 8080")

	// start server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println(debug.Stack())
		return
	}
}
