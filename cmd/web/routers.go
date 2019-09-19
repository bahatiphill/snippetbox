package main

import "net/http"
import "github.com/justinas/alice"

func (app *application) routes() http.Handler {

	//Create a middelware chain 
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// wrap the exsisting chain witrh the recoverPanic middleware
	return standardMiddleware.Then(mux)
}