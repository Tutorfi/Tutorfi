package storage

import (
	"app/internal/models"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)
func RowToAccount(row *sql.Row) (models.Account){
	var acc models.Account
	row.Scan(&acc.Id, &acc.Firstname, &acc.Lastname, &acc.Email, &acc.Password, &acc.SessionId)
	fmt.Println(acc)
	return acc
}
func (s *PostgresStorage) GetAccount(email string) (*models.Account, error) {
	var acc models.Account
	res := s.db.QueryRow("SELECT id, firstname, lastname, email, password, sessionid FROM account WHERE email = $1", email)
	if res.Scan(&acc.Id) != nil{
		return nil, sql.ErrNoRows
	}
	if res.Err() != nil{
		return nil, res.Err()
	}
	acc = RowToAccount(res)
	return &acc, nil
}
//Inserts an account into the database, does not return the created account.
func (s *PostgresStorage) CreateAccount(fname, lname, email, password string) (*models.Account, error){
	var acc models.Account
	res := s.db.QueryRow("INSERT INTO account VALUES (DEFAULT, null, $1, $2, $3, $4)", fname, lname, email, password)
	if res.Scan(&acc.Id) != nil{
		return nil, sql.ErrNoRows
	}
	if res.Err() != nil{
		return nil, res.Err()
	}
	acc = RowToAccount(res)
	return &acc, nil
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

