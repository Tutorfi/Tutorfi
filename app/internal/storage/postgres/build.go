package storage

import "fmt"
import "golang.org/x/crypto/bcrypt"

func (s *PostgresStorage) BuildDevDB() error {
	val := ` 
    DROP TABLE IF EXISTS "group";
    DROP TABLE IF EXISTS "schedule";
    DROP TABLE IF EXISTS "user_schedule";
    DROP TABLE IF EXISTS "permission";
    DROP TABLE IF EXISTS "account";
    DROP TABLE IF EXISTS "organization";
    
    CREATE TABLE "organization" (
        "id" integer UNIQUE PRIMARY KEY,
        "setting" varchar
    );

    CREATE TABLE "account" (
      "id" uuid UNIQUE PRIMARY KEY DEFAULT (gen_random_uuid()),
      "seesion_id" uuid UNIQUE,
      "organization_id" integer UNIQUE,
      "email" varchar UNIQUE NOT NULL,
      "firstname" varchar NOT NULL,
      "lastname" varchar NOT NULL,
      "password" varchar NOT NULL,
      "created_at" timestamp,
      FOREIGN KEY ("organization_id") REFERENCES "organization"("id")
    );
    
    CREATE TABLE "group" (
      "id" integer UNIQUE PRIMARY KEY,
      "organization_id" integer,
      "name" varchar UNIQUE NOT NULL,
      "data" jsonb,
      FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE
    );
    
    CREATE TABLE "user_schedule" (
      "id" integer UNIQUE PRIMARY KEY,
      "account_id" uuid UNIQUE NOT NULL,
      "data" jsonb,
      FOREIGN KEY ("account_id") REFERENCES "account"("id") ON DELETE CASCADE
    );
    
    CREATE TABLE "permission" (
      "id" integer UNIQUE PRIMARY KEY,
      "account_id" uuid UNIQUE NOT NULL,
      "permissions" jsonb,
      FOREIGN KEY ("account_id") REFERENCES "account"("id") ON DELETE CASCADE
    );
    `
	_, err := s.db.Exec(val)
	if err != nil {
		fmt.Println("unable to create database")
        fmt.Println(err)
        return err;
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte("passwordthing"), 0)
	_, err = s.db.Exec("INSERT INTO account (firstname,lastname,email,password) VALUES ('bob', 'Builder', 'bob@gmail.com', $1)", hash)
	if err != nil {
		fmt.Println("unable to insert values into test database")
		fmt.Println(err)
        return err;
	}
	hash, _ = bcrypt.GenerateFromPassword([]byte("passwordthing"), 0)
	_, err = s.db.Exec("INSERT INTO account (firstname,lastname,email,password) VALUES ('Jane', 'Lin', 'JaneLin@gmail.com', $1)", hash)
	if err != nil {
		fmt.Println("unable to insert values into test database")
		fmt.Println(err)
        return err;
	}
	hash, _ = bcrypt.GenerateFromPassword([]byte("passwordthing"), 0)
	_, err = s.db.Exec("INSERT INTO account (firstname,lastname,email,password) VALUES ('Me', 'Bulmaro', 'Bulmaro@gmail.com', $1)", hash)
	if err != nil {
		fmt.Println("unable to insert values into test database")
		fmt.Println(err)
        return err;
	}
	hash, _ = bcrypt.GenerateFromPassword([]byte("passwordthing"), 0)
	_, err = s.db.Exec("INSERT INTO account (firstname,lastname,email,password) VALUES ('John', 'Doe', 'JohnDoe@gmail.com', $1)", hash)
	if err != nil {
		fmt.Println("unable to insert values into test database")
		fmt.Println(err)
        return err;
	}
	fmt.Println("Finished building db")
    return nil;
}
