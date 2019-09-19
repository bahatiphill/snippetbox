package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

//ServerError helper writes an error and stack trace to the errorLog
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

//clientError helper send a specific status code and corresponding descirption to the user
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// NotFound helper
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {

	//retrieve the appropriate template set from the cache based on the page name
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}
	//Execute the template set, passing in any dynamic data
	err := ts.Execute(w, td)
	if err != nil {
		app.serverError(w, err)
	}
}
