package app

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
	"fmt"
)

func ConnectPgsql() (*sql.DB, error) {

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)
	fmt.Println(psqlconn)
	db, err := sql.Open("pgx", psqlconn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	return db, err
}

func ConnectPgsqlTest() (*sql.DB, error){
	psqlconn := fmt.Sprintf("host=pgsqlTest user=user password=postgres dbname=master")
	db, err := sql.Open("pgx", psqlconn)
	if err != nil{
		return nil, err
	}
	err = db.Ping()

	return db, err
}
