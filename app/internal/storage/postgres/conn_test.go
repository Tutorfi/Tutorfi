package storage

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"testing"
)

func TestConn(t *testing.T) {
	psqlconn := ("host=pgsqlTest user=user password=postgres dbname=master sslmode=disable port=5432")
	db, err := sql.Open("pgx", psqlconn)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Printf(err.Error())
	}

}
