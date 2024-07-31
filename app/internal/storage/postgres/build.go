package storage

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s *PostgresStorage) BuildDevDB() error {
	val := ` 
    DROP TABLE IF EXISTS "group_account";
    DROP TABLE IF EXISTS "group";
    DROP TABLE IF EXISTS "schedule";
    DROP TABLE IF EXISTS "user_schedule";
    DROP TABLE IF EXISTS "permission";
    DROP TABLE IF EXISTS "tag_account";
	DROP TABLE IF EXISTS "tag";
    DROP TABLE IF EXISTS "account";
    DROP TABLE IF EXISTS "organization";

    CREATE TABLE "organization" (
        "id" SERIAL UNIQUE PRIMARY KEY,
        "name" varchar UNIQUE NOT NULL
    );

    CREATE TABLE "account" (
      "id" uuid UNIQUE PRIMARY KEY DEFAULT (gen_random_uuid()),
      "session_id" uuid UNIQUE,
      "organization_id" INTEGER UNIQUE,
      "email" varchar UNIQUE NOT NULL,
      "username" varchar NOT NULL,
      "firstname" varchar NOT NULL,
      "lastname" varchar NOT NULL,
      "password" varchar NOT NULL,
      "created_at" timestamptz DEFAULT current_timestamp,
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
      "organization_id" INTEGER NOT NULL,
      "name" varchar UNIQUE NOT NULL,
      "data" jsonb,
      FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE
    );

    CREATE TABLE "group_account" (
	    "id" SERIAL UNIQUE PRIMARY KEY,
	    "group_id" integer,
	    "account_id" uuid,
	    FOREIGN KEY ("group_id") REFERENCES "group"("id") ON DELETE CASCADE,
	    FOREIGN KEY ("account_id") REFERENCES "account"("id") ON DELETE CASCADE
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
    `
	_, err := s.db.Exec(val)
	if err != nil {
		fmt.Println("unable to create database")
		fmt.Println(err)
		return err
	}
	fmt.Println("Created the tables")
	_, err = s.db.Exec(`INSERT INTO "organization" ("name") VALUES ('Tutorfi')`)
	if err != nil {
		fmt.Println("unable to insert values into test database")
		fmt.Println(err)
		return err
	}
	users := [][]string{
        {"bob", "Builder", "BuilderB5", "bob@gmail.com"},
		{"Jane", "Lin", "LinJ3", "Jane@gmail.com"},
		{"Me", "Bull", "BullM", "Bull@gmail.com"},
		{"John", "Doe", "DoeJ", "JohnDoe@gmail.com"},
	}
	var orgId int
	err = s.db.QueryRow(`SELECT ("id") FROM "organization" WHERE "name"='Tutorfi'`).Scan(&orgId)
	if err != nil {
		fmt.Println("Unable to read organization row")
		fmt.Println(err)
		return err
	}
	_, err = s.db.Exec(`INSERT INTO "group" ("organization_id","name") VALUES ($1, 'Linear Algebra')`, orgId)
	if err != nil {
		fmt.Println("unable to insert users into database")
		fmt.Println(err)
		return err
	}
	for i := 0; i < len(users); i++ {
		hash, _ := bcrypt.GenerateFromPassword([]byte("password"), 0)
		_, err = s.db.Exec(`INSERT INTO "account" (firstname,lastname,email,username,password) 
        VALUES ($1, $2, $3, $4, $5)`, users[i][0], users[i][1],
			users[i][2], users[i][3], hash)
		if err != nil {
			fmt.Println("unable to insert users into database")
			fmt.Println(err)
			return err
		}
		var userId string
		err = s.db.QueryRow(`SELECT ("id") FROM "account" WHERE "email"=$1`, users[i][2]).Scan(&userId)
		if err != nil {
			fmt.Println("Unable to read user row")
			fmt.Println(err)
			return err
		}
		_, err = s.db.Exec(`INSERT INTO "group_account" ("group_id", "account_id") VALUES ($1, $2)`, orgId, userId)
		if err != nil {
			fmt.Println("unable to insert users into database")
			fmt.Println(err)
			return err
		}
	}
	fmt.Println("Finished building db")
	return nil
}
