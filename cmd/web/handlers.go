package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

var pathToTemplate = "./templates"

const (
	pathToPage   = "page"
	pathToLayout = "layout"
)

// TemplateData template data custom type
type TemplateData struct {
	IP   string         //IP ADDRESS
	Data map[string]any // DATA map
}

// HandlerHome -> Home handler
func (app *application) HandlerHome(w http.ResponseWriter, r *http.Request) {
	var td = make(map[string]any)
	// check if session key exist
	if app.Session.Exists(r.Context(), "test") {
		td["test"] = app.Session.GetString(r.Context(), "test")
	} else {
		app.Session.Put(r.Context(), "test", "Hit this page at"+time.Now().String())
	}

	_ = app.Render(w, path.Join(pathToPage, "home.page.gohtml"), &TemplateData{
		Data: td,
		// add IP address
		IP: app.ipFromContext(r.Context()),
	})
}

// Render page
func (app *application) Render(w http.ResponseWriter, t string, data *TemplateData) error {
	// file to render on each page
	files := []string{
		path.Join(pathToTemplate, t),
		path.Join(pathToTemplate, pathToLayout, "base.layout.gohtml"),
		path.Join(pathToTemplate, pathToPage, "login.page.gohtml"),
	}
	// parse the template from disk
	parsedTemp, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	// Execute template
	if err := parsedTemp.Execute(w, data); err != nil {
		return err
	}

	return nil
}

// HandlerLogin -> Handle login page events
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
