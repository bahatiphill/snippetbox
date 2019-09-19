package main

import (
	"html/template"
	"path/filepath"
	"snippetbox-modules/pkg/models"
	"time"
	"net/url"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	FormData	url.Values
	FormErrors	map[string]string

}

// Function which return a human readablr time.Time object
func humanDate(t time.Time) string {
	return t.Format("03 Jan 2009 at 15:57")
}

//initialize a template.FuncMap which hold custom template functions
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}
	//slice of all filepaths with the extension ".page.tmpl"
	pages, err := filepath.Glob(filepath.Join(dir, "*page.tmpl"))
	if err != nil {
		return nil, err
	}
	//Loop through the pages one-by-one
	for _, page := range pages {

		//Get the file name from the full path and assign it to the name variable
		name := filepath.Base(page)

		//Register humanDate fuctions using Funcs method then parse the page template in to a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		//Use the ParseGlob method to add any 'layout' templates to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*layout.tmpl"))
		if err != nil {
			return nil, err
		}
		//Use the ParseGlob method to add any 'partial' templates to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*partial.tmpl"))
		if err != nil {
			return nil, err
		}

		//Add the template set to the cache
		cache[name] = ts
	}
	// return the map
	return cache, nil
}
