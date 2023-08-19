package main

import (
	"database/sql"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// ping db and check connectivity
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func (app *application) connectToDB() (*sql.DB, error) {
	return openDB(app.DSN)
}
