package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
	DSN     string
	DB      *sql.DB
}

func main() {
	// init app config
	app := application{
		Session: getSession(),
	}

	// flag variable
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "POSTGRES CONNECTION STRING")
	flag.Parse()

	// connect to db
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	if conn != nil {
		log.Println("CONNECTED TO DB")
	}

	// assign conn to struct variable
	app.DB = conn

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
