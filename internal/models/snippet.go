package models

import (
	"database/sql"
	"errors"
	"time"
)

// snippet data model
type Snippet struct {
	ID        int
	Title     string
	Content   string
	Expires   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// get single snippet by id
func (m *SnippetModel) Get(id int) (Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	var s Snippet

	if err := row.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.Expires); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil
}

// insert or save new snippet

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

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

// get lates snippet list
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return []Snippet{}, nil
}
