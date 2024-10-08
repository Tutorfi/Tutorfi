package storage

import (
	"app/internal/models"
	"database/sql"
	_ "database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func (s *PostgresStorage) GetAccount(email string) (*models.Account, error) {
	var acc models.Account
	err := s.db.QueryRow("SELECT id, firstname, lastname, email, username, password, session_id, organization_id FROM \"account\" WHERE email = $1", email).Scan(
		&acc.Id, &acc.Firstname, &acc.Lastname, &acc.Email, &acc.Username, &acc.Password, &acc.SessionId, &acc.OrganizationId)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

func (s *PostgresStorage) GetAccountSessionId(sessionId string) (*models.Account, error) {
	var acc models.Account
	err := s.db.QueryRow(`
        SELECT id, firstname, lastname, email, username, password, session_id, 
        organization_id FROM "account" WHERE "session_id" = $1`, sessionId).Scan(
		&acc.Id, &acc.Firstname, &acc.Lastname, &acc.Email, &acc.Username, &acc.Password, &acc.SessionId,
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
func (s *PostgresStorage) CreateAccount(fname string, lname string, email string, username string, password string) error {
	//var acc models.Account
    res := s.db.QueryRow(`INSERT INTO "account" ("firstname", "lastname", "email", "username", "password") VALUES 
                        ($1, $2, $3, $4, $5)`, fname, lname, email, username, password)
	return res.Err()
}

func (s *PostgresStorage) GetPassword(email string) (string, error) {
	var password string
	err := s.db.QueryRow(`SELECT "password" FROM "account" WHERE "email" = $1`, email).Scan(password)
	return password, err
}

func (s *PostgresStorage) UpdateEmail(id string, email string) error {
    _, err := s.db.Exec(`UPDATE "account" SET "email" = $1 WHERE "id" = $2`, email, id);
    return err
}

func (s *PostgresStorage) UpdateUsername(id string, username string) error {
    _, err := s.db.Exec(`UPDATE "account" SET "username" = $1 WHERE "id" = $2`, username, id);
    return err
}

func (s *PostgresStorage) SetSessionID(id string, sessionid string) error {
	_, err := s.db.Exec(`UPDATE "account" SET "session_id" = $1 WHERE "id" = $2`, sessionid, id)
	return err
}

func (s *PostgresStorage) DeleteAccount(id string) error {
	num, err := s.db.Exec(`DELETE FROM "account" WHERE "id" = $1`, id)
	fmt.Println(num)
	return err
}

func (s *PostgresStorage) ResetSessionID(sessionid string) error {
	_, err := s.db.Exec(`UPDATE "account" SET "session_id" = null WHERE "session_id" = $1`, sessionid)
	return err
}
