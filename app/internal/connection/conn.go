package connection

import (
	"fmt"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)

// Change the name of this function to something more descriptive
func ConnectPgsql() (*sql.DB, error) {

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)
	
	return sql.Open("postgres", psqlconn)
}

