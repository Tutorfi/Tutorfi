package storage

import (
	"app/internal/models"
	"database/sql"
	_ "database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func (s *PostgresStorage) GetAccount(email string) (*models.Account, error) {
	var acc models.Account
	err := s.db.QueryRow("SELECT id, firstname, lastname, email, password, session_id, organization_id FROM \"account\" WHERE email = $1", email).Scan(
		&acc.Id, &acc.Firstname, &acc.Lastname, &acc.Email, &acc.Password, &acc.SessionId, &acc.OrganizationId)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

func (s *PostgresStorage) GetAccountSessionId(c echo.Context) (*models.Account, error) {
	var acc models.Account
	cookie, err := c.Cookie("Tutorfi_Account")
	if err != nil {
		return &models.Account{}, err
	}
    sessionId := cookie.Value
	err = s.db.QueryRow(`
        SELECT id, firstname, lastname, email, password, session_id, 
        organization_id FROM "account" WHERE "session_id" = $1`, sessionId).Scan(
		    &acc.Id, &acc.Firstname, &acc.Lastname, &acc.Email, &acc.Password, &acc.SessionId, 
            &acc.OrganizationId)
	if err == sql.ErrNoRows {
		return &models.Account{}, nil
	}
	if err != nil {
		fmt.Println("A database Error Occured")
		fmt.Println(err)
		return &models.Account{}, err
	}
	return &acc, nil
}

// Inserts an account into the database, does not return the created account.
func (s *PostgresStorage) CreateAccount(fname, lname, email, password string) error {
	//var acc models.Account
	res := s.db.QueryRow("INSERT INTO \"account\" VALUES (DEFAULT, null, null, $1, $2, $3, $4, null)", fname, lname, email, password)
	return res.Err()
}

func (s *PostgresStorage) GetPassword(email string) (string, error) {
	var password string
	err := s.db.QueryRow("SELECT password FROM \"account\" WHERE email = $1", email).Scan(password)
	return password, err
}

func (s *PostgresStorage) SetSessionID(email string, sessionid string) error {
	_, err := s.db.Exec("UPDATE \"account\" SET session_id = $1 WHERE email = $2", sessionid, email)
	return err
}

func (s *PostgresStorage) DeleteAccount(id string) error {
	num, err := s.db.Exec(`DELETE FROM "account" WHERE id = $1`, id)
	fmt.Println(num)
	return err
}

func (s *PostgresStorage) ResetSessionID(id string) error {
	_, err := s.db.Exec(`UPDATE "account" SET session_id = null WHERE id = $1`, id)
	return err
}
