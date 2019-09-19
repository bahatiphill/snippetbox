package main

import (
	"fmt"
	"net/http"
	"snippetbox-modules/pkg/models"
	"strconv"
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
	w.Write([]byte("Create a new snippet.."))
}

//Create a Snippet handler
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	// Dummy data
	title := "Dummy title"
	content := "Dummy content"
	expires := "7"

	// pass the data to the SnippetModel.Insert, return the ID of the new record
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	//Redirect the user to the relevant page for the snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
