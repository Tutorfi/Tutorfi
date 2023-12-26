package connection

import (
	"fmt"
	"os"
	"database/sql"
	"github.com/joho/godotenv"
)

// Change the name of this function to something more descriptive
func ConnectPgsql() (*sql.DB, error) {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	psqlconn := fmt.Sprintf("host=%s port=5432 user=app password=%s dbname=%s sslmode=disable", host, password, dbName)
	
	return sql.Open("postgres", psqlconn)
}

