package postgres

import (
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"sa.com/pkg/models"
)

type SnippetModel struct {
	Pool *pgxpool.Pool
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	sqlStatement := "INSERT INTO snippets (title, content, created, expires) VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + ($3 * INTERVAL '365 DAY')) returning id"
	id := 0
	err := m.Pool.QueryRow(sqlStatement, title, content, expires).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > NOW() AND id = $1`

	row := m.Pool.QueryRow(stmt, string(id))

	s := &models.Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
