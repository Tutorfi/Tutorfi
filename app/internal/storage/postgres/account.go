package storage

import (
	"app/internal/models"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"reflect"
)
func RowToAccount(row *sql.Rows) (models.Account){
	
	var acc models.Account
	return acc
}
func (s *PostgresStorage) GetAccount(email string) (*models.Account, error) {
	var acc models.Account
	res, err := s.db.Query("SELECT id, firstname, lastname, email, password, sessionid FROM account WHERE email = $1", email)
	
	fmt.Println(reflect.TypeOf(res))
	fmt.Println(res)
	fmt.Println(err)
	if err == sql.ErrNoRows{
		return nil, err
	}
	acc = RowToAccount(res)
	return &acc, err
}
//Inserts an account into the database, does not return the created account.
func (s *PostgresStorage) CreateAccount(fname, lname, email, password string) (error){
	_, err := s.db.Exec("INSERT INTO account VALUES (DEFAULT, null, $1, $2, $3, $4)", fname, lname, email, password)
	return err
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
	_, err := s.db.Exec(`DELETE FROM account WHERE id = $1`, id)
	return err
}

func (s *PostgresStorage) ResetSessionID(id string) (error){
	_, err := s.db.Exec("UPDATE account SET sessionid = null WHERE id = $1", id)
	return err
}

