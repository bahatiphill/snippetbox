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
	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES (?,?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Get a specific snippet based in its ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id) 

	s := &models.Snippet{}

	// Use row.scan to copy values from each field in sql.row to convert them into corresponding Go datatype in Snippet struct
	err :=  row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}

// Get 10 most recent snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt :=  `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP ORDER BY created LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize a empty slice to hold the models.Snippets struct
	snippets := []*models.Snippet{}

	// Use rows.Next() to iterate through the returned result
	for rows.Next() {
		//create a pointer to initialize  a new Snippet struct
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		//Append it to the slice of snippets
		snippets = append(snippets, s)
	}
	//check if the iteration thru rows haven't counted any error
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// If everthing went OK, return the Snippets slice
	return snippets, nil
}