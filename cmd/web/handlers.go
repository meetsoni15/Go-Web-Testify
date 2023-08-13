package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

var pathToTemplate = "./templates"

type TemplateData struct {
	IP   string         //IP ADDRESS
	Data map[string]any // DATA map
}

func (app *application) HandlerHome(w http.ResponseWriter, r *http.Request) {
	_ = app.Render(w, r, "home.page.gohtml", &TemplateData{})
}

func (app *application) Render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	// parse the template from disk
	parsedTemp, err := template.ParseFiles(
		path.Join(pathToTemplate, t),
		path.Join(pathToTemplate, "login.page.gohtml"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	// add IP address
	data.IP = app.ipFromContext(r.Context())

	// Execute template
	if err := parsedTemp.Execute(w, data); err != nil {
		return err
	}

	return nil
}

func (app *application) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	// parse form
	if err := r.ParseForm(); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadGateway)
		return
	}

	// extract email and password
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	log.Println(email, password)

	fmt.Fprint(w, email)
}
