package main

import (
	"log"
	"testing"
	_ "github.com/jackc/pgx/v5"
	"app/internal/app"
)

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	log.Println("works")
	_, err := app.BuildTestDB()
	if err != nil{
		log.Println("stuff")
	}
}

