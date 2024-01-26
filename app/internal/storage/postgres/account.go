package storage

import (
	"app/internal/models"
	_ "database/sql"
	_ "github.com/lib/pq"
	"fmt"
)
func (s *PostgresStorage) GetAccount(email string) (*models.Account, error) {
	var acc models.Account
	err := s.db.QueryRow("SELECT id, firstname, lastname, email, password, seesion_id, organization_id FROM account WHERE email = $1", email).Scan(&acc.Id, &acc.Firstname, &acc.Lastname, &acc.Email, &acc.Password, &acc.SessionId, &acc.OrganizationId)
	if err != nil{
		return nil, err
	}
	return &acc, nil
}
//Inserts an account into the database, does not return the created account.
func (s *PostgresStorage) CreateAccount(fname, lname, email, password string) (error){
	//var acc models.Account
	res := s.db.QueryRow("INSERT INTO account VALUES (DEFAULT, null, $1, $2, $3, $4)", fname, lname, email, password)
	return res.Err()
}

func (s *PostgresStorage) GetPassword(email string) (string, error) {
	var password string
	err := s.db.QueryRow("SELECT password FROM account WHERE email = $1", email).Scan(password)
	return password, err
}

func (s *PostgresStorage) SetSessionID(email string, sessionid string) (error){
	_, err := s.db.Exec("UPDATE account SET sessionid = $1 WHERE email = $2", sessionid, email)
	return err
}

func (s *PostgresStorage) DeleteAccount(id string) (error){
	num, err := s.db.Exec(`DELETE FROM account WHERE id = $1`, id)
	fmt.Println(num)
	return err
}

func (s *PostgresStorage) ResetSessionID(id string) (error){
	_, err := s.db.Exec("UPDATE account SET sessionid = null WHERE id = $1", id)
	return err
}

