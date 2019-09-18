package main

import "snippetbox-modules/pkg/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}