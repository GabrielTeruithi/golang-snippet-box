package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, createdAt, expiresAt string) error {
	stmt := `INSERT INTO snippets (title, content, created,expires) VALUES ($1, $2, NOW(), CURRENT_DATE + INTERVAL $4 day)`

	_, err := m.DB.Exec(stmt, title, content, createdAt, expiresAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *SnippetModel) et(id int) (Snippet, error) {
	return Snippet{}, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
