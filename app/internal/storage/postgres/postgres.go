package storage

import (
	"app/internal/models"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
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
	s.db.QueryRow("SELECT * FROM account WHERE account.Email = $1", email)
	return &models.Account{}, nil
}
//Insert a new account into the database, returning the account that was just created
func (s *PostgresStorage) CreateAccount(fname, lname, email, password string) (error){
	_, err := s.db.Exec("INSERT INTO account VALUES (0, $1, $2, $3, $4)", fname, lname, email, password)
	return err
}

func (s *PostgresStorage) GetPassword(email string) (string, error) {
	var password string
	err := s.db.QueryRow("SELECT Password FROM account WHERE account.Email = $1", email).Scan(password)
	return password, err
}