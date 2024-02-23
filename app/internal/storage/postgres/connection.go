package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
)

func ConnectPgsql() (*sql.DB, error) {

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)
	db, err := sql.Open("pgx", psqlconn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	return db, err
}

func ConnectPgsqlTest() (*sql.DB, error) {
	psqlconn := ("host=pgsqlTest user=user password=postgres dbname=master sslmode=disable port=5432")
	db, err := sql.Open("pgx", psqlconn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	return db, err
}
