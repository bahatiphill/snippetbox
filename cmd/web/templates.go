package main

import (
	"html/template"
	"path/filepath"
	"snippetbox-modules/pkg/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
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

		//parse the page template in to a template set
		ts, err := template.ParseFiles(page)
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
