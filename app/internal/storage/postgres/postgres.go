package storage

import (
	"app/internal/models"
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		db: db,
	}
}

func (s *PostgresStorage) GetAccount(email string) (*models.Account, error) {
	s.db.QueryRow("SELECT * FROM accounts WHERE username = $1", email)

	return &models.Account{}, nil
}