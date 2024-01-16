package storage

import "fmt"
import "golang.org/x/crypto/bcrypt"
func (s *PostgresStorage) BuildDevDB() {
	val := `
	DROP TABLE IF EXISTS account;
	CREATE TABLE account (
	id uuid DEFAULT gen_random_uuid (),
	sessionid uuid,
	firstname TEXT NOT NULL,
	lastname TEXT NOT NULL,
	email TEXT NOT NULL,
	password VARCHAR NOT NULL,
	PRIMARY KEY(ID)
	);
	`
	_, err := s.db.Exec(val)
	if err != nil {
		fmt.Println("unable to create database")
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte ("passwordthing"), 0)
	_, err = s.db.Exec("INSERT INTO account (firstname,lastname,email,password) VALUES ('bob', 'Builder', 'bob@gmail.com', $1)", hash)
	if err != nil {
		fmt.Println("unable to insert values into test database")
		fmt.Println(err)
	}
	
}
