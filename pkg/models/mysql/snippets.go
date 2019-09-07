package mysql

import (
	"database/sql"

	"github.com/leandroxt/snippetbox/pkg/models"
)

// SnippetModel Define a type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := "INSERT INTO snippets (title, content, created, expires) VALUES ($1, $2, now(), now() + INTERVAL '10' DAYS) RETURNING id"

	var lastInsertID int

	// We cannot user DB.Exec() and user LastInsertId() because postgres driver does not support it.
	err := m.DB.QueryRow(stmt, title).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// Get a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
