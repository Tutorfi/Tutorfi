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

func (s *PostgresStorage) CreateAccount(fname, lname, email, password string) (*models.Account, error){
	s.db.Query("INSERT INTO users VALUES (0, $1, $2, $3, $4)", fname, lname, email, password)
	return &models.Account{}, nil
}

func (s *PostgresStorage) GetPassword(email string) (*models.Account, error) {
	s.db.QueryRow("SELECT password FROM accounts WHERE username = $1", email)

	return &models.Account{}, nil
}