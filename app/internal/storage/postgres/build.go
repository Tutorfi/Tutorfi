package storage

import "fmt"

func (s *PostgresStorage) BuildDevDB() {
	val := `
	DROP TABLE IF EXISTS account;
	CREATE TABLE account (
	id uuid DEFAULT gen_random_uuid (),
	altid uuid,
	firstname TEXT NOT NULL,
	lastname TEXT NOT NULL,
	email TEXT NOT NULL,
	password VARCHAR NOT NULL,
	PRIMARY KEY(ID)
	);
	`
	_, err := s.db.Exec(val)
	if err != nil {
		fmt.Println("unable to create table")
	}
	_, err = s.db.Exec("INSERT INTO account (firstname,lastname,email,password) VALUES ('bob', 'Builder', 'bob@gmail.com', 'passwordthing');")
	if err != nil {
		fmt.Println("unable to insert values")
		fmt.Println(err)
	}
	
}
