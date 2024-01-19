package storage

import "fmt"
import "golang.org/x/crypto/bcrypt"

func (s *PostgresStorage) BuildDevDB() {
	val := ` 
    DROP TABLE IF EXISTS schedule;
    DROP TABLE IF EXISTS user_schedule;
    DROP TABLE IF EXISTS permission;
	DROP TABLE IF EXISTS account;
    DROP TABLE IF EXISTS organization;
    
    CREATE TABLE organization (
      id integer PRIMARY KEY UNIQUE,
      setting text
    );
    
    CREATE TABLE account (
      id uuid DEFAULT gen_random_uuid () PRIMARY KEY UNIQUE,
      organization_id integer UNIQUE,
      email varchar UNIQUE,
      firstname varchar,
      lastname varchar,
      password varchar,
      created_at timestamp
    );
    
    CREATE TABLE schedule (
      id integer PRIMARY KEY UNIQUE,
      organization_id integer UNIQUE
    );
    
    CREATE TABLE user_schedule (
      id integer PRIMARY KEY UNIQUE,
      account_id uuid UNIQUE,
      data jsonb
    );
    
    CREATE TABLE permission (
      id integer PRIMARY KEY UNIQUE,
      account_id uuid UNIQUE,
      permissions jsonb
    );
    
    ALTER TABLE account ADD FOREIGN KEY (organization_id) REFERENCES organization (id);
	
    ALTER TABLE user_schedule ADD FOREIGN KEY (account_id) REFERENCES account (id)
	on delete cascade on update cascade;
    
	ALTER TABLE schedule ADD FOREIGN KEY (organization_id) REFERENCES organization (id)
	on delete cascade on update cascade;
	
    ALTER TABLE permission ADD FOREIGN KEY (account_id) REFERENCES account (id)
	on delete cascade on update cascade;
	
	ALTER TABLE account ALTER COLUMN organization_id DROP NOT NULL;
	ALTER TABLE user_schedule ALTER COLUMN account_id DROP NOT NULL;
	ALTER TABLE permission ALTER COLUMN account_id DROP NOT NULL;
    `
	_, err := s.db.Exec(val)
	if err != nil {
		fmt.Println("unable to create table")
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte("passwordthing"), 0)
	_, err = s.db.Exec("INSERT INTO account (firstname,lastname,email,password) VALUES ('bob', 'Builder', 'bob@gmail.com', $1)", hash)
	if err != nil {
		fmt.Println("unable to insert values")
		fmt.Println(err)
	}
    fmt.Println("Finished building db")
}
