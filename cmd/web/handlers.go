package main

import (
	"fmt"
	"net/http"
	"snippetbox-modules/pkg/models"
	"snippetbox-modules/pkg/forms"
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
	app.render(w, r, "create.page.tmpl", &templateData{
		//Pass an empty forms.Form object to the template
		Form: forms.New(nil),
	})
}

//Create a Snippet handler
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	//use r.PostForm.Get to retrieve the data
	//create a new forms.Form struct containing the posted DATA and validate them 
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")



	

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{
			Form: form,
		})
		return
	}

	// pass the data to the SnippetModel.Insert, return the ID of the new record
	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	//Redirect the user to the relevant page for the snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
