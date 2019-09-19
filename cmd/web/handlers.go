package main

import (
	"fmt"
	"net/http"
	"snippetbox-modules/pkg/models"
	"strconv"
	"strings"
	"unicode/utf8"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	//Because Pat matches the "/" path exactly, we cna now remove the manual check
	// if r.URL.Path != "/" {
	// 	app.notFound(w)
	// 	return
	// }
	
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{Snippets: s})
}

// Displaying Specific Snippet Handler
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "show.page.tmpl", &templateData{Snippet: s})
}

//createSnippetForm handler
func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}

//Create a Snippet handler
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	//use  r.PostForm.Get to retrieve the data
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	//Initialize a map to hold any validation erros
	errors := make(map[string]string)

	//***************** DATA VALIDATION ***********

	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This field is too long (only 100 characters)"
	}

	if strings.TrimSpace(content) == "" {
		errors["content"] = "This field cannot be empty"
	}
	if strings.TrimSpace(expires) == "" {
		errors["expires"] = "This cannot be empty"
	} else if expires != "365" && expires != "7" && expires != "1" {
		errors["expires"] = "This field is invalid"
	}
	//****************************************

	if len(errors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: errors,
			FormData: r.PostForm,
		})
		return
	}

	// pass the data to the SnippetModel.Insert, return the ID of the new record
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	//Redirect the user to the relevant page for the snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
