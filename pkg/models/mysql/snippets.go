package mysql 

import (
	"database/sql"
	"snippetbox-modules/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// Create a new snippet
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get a specific snippet based in its ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Get 10 most recent snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}