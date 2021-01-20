package postgres

import (
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
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
