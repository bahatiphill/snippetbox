package main

import "net/http"
import "github.com/justinas/alice"
import "github.com/bmizerany/pat"

func (app *application) routes() http.Handler {

	//Create a middelware chain 
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	//mux := http.NewServeMux()
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// wrap the exsisting chain witrh the recoverPanic middleware
	return standardMiddleware.Then(mux)
}