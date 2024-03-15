package storage

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s *PostgresStorage) BuildDevDB() error {
	val := ` 
  DROP TABLE IF EXISTS "user_schedule";
  DROP TABLE IF EXISTS "files_group";
  DROP TABLE IF EXISTS "schedule";
  DROP TABLE IF EXISTS "permission";
  DROP TABLE IF EXISTS "files";
  DROP TABLE IF EXISTS "group";
  DROP TABLE IF EXISTS "account";
  DROP TABLE IF EXISTS "organization";
  
  CREATE TABLE "organization" (
      "id" integer UNIQUE PRIMARY KEY,
      "setting" varchar
  );

  CREATE TABLE "materials"(
    "id" SERIAL UNIQUE PRIMARY KEY,
    "author" SERIAL,
    "owner" SERIAL,
    "group" SERIAL,
    "data" jsonb
  );
  CREATE TABLE ""(

  );
    CREATE TABLE "account" (
      "id" uuid UNIQUE PRIMARY KEY DEFAULT (gen_random_uuid()),
      "session_id" uuid UNIQUE,
      "organization_id" INTEGER UNIQUE,
      "email" varchar UNIQUE NOT NULL,
      "firstname" varchar NOT NULL,
      "lastname" varchar NOT NULL,
      "password" varchar NOT NULL,
      "created_at" timestamp,
      FOREIGN KEY ("organization_id") REFERENCES "organization"("id")
    );

    CREATE TABLE "tag" (
	  "id" SERIAL UNIQUE PRIMARY KEY,
	  "name" varchar UNIQUE NOT NULL,
      "delete_group" BOOLEAN DEFAULT FALSE,
      "remove_add_student" BOOLEAN DEFAULT FALSE,
      "remove_add_access_level" BOOLEAN DEFAULT FALSE
    );

    CREATE TABLE "tag_account" (
      "id" SERIAL UNIQUE PRIMARY KEY,
      "account_id" uuid,
      "tag_id" INT,
      FOREIGN KEY ("account_id") REFERENCES "account"("id") ON DELETE CASCADE,
      FOREIGN KEY ("tag_id") REFERENCES "tag"("id") ON DELETE CASCADE
    );

    CREATE TABLE "group" (
      "id" SERIAL UNIQUE PRIMARY KEY,
      "organization_id" INTEGER,
      "name" varchar UNIQUE NOT NULL,
      "data" jsonb,
      FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE
    );
    
    CREATE TABLE "user_schedule" (
      "id" SERIAL UNIQUE PRIMARY KEY,
      "account_id" uuid UNIQUE NOT NULL,
      "data" jsonb,
      FOREIGN KEY ("account_id") REFERENCES "account"("id") ON DELETE CASCADE
    );

    CREATE TABLE "permission" (
      "id" SERIAL UNIQUE PRIMARY KEY,
      "account_id" uuid UNIQUE NOT NULL,
      "permissions" jsonb,
      FOREIGN KEY ("account_id") REFERENCES "account"("id") ON DELETE CASCADE
    );
    CREATE TABLE "files" (
      "id" integer UNIQUE PRIMARY KEY,
      "account_id" uuid,
      "filedir" varchar,
      FOREIGN KEY ("account_id") REFERENCES "account"("id") ON DELETE CASCADE
    );
    CREATE TABLE "files_group" (
      "id" integer UNIQUE PRIMARY KEY,
      "files_id" integer,
      "group_id" integer,
      FOREIGN KEY ("files_id") REFERENCES "files"("id") ON DELETE CASCADE,
      FOREIGN KEY ("group_id") REFERENCES "group"("id") ON DELETE CASCADE
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
