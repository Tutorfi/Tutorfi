package storage

import (
	"app/internal/models"
	"database/sql"

	"github.com/google/uuid"
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
	var acc models.Account
	err := s.db.QueryRow("SELECT id, altid, fname, lname, email, password FROM account WHERE email = $1", email).Scan(&acc.ID, &acc.AltID, &acc.Firstname, &acc.Lastname, &acc.Email, &acc.Password)
	if err == sql.ErrNoRows{
		return nil, err
	}
	return &acc, err
}
//Insert a new account into the database, returning the account that was just created
func (s *PostgresStorage) CreateAccount(fname, lname, email, password string) (error){
	_, err := s.db.Exec("INSERT INTO account VALUES ($1, $2, $3, $4, $5)",uuid.NewString(), fname, lname, email, password)
	return err
}

func (s *PostgresStorage) GetPassword(email string) (string, error) {
	var password string
	err := s.db.QueryRow("SELECT password FROM account WHERE email = $1", email).Scan(password)
	return password, err
}
