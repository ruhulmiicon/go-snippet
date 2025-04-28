package models

import (
	"database/sql"
	"time"
)

// snippet data model
type Snippet struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// get single snippet by id
func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// insert or save new snippet

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// get lates snippet list
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return []Snippet{}, nil
}
