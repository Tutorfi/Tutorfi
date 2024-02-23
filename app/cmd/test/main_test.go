package main

import (
	"app/internal/app"
	_ "github.com/jackc/pgx/v5"
	"testing"
	"log"
)

func TestMain(t *testing.M) {
	err := app.BuildTestDB()
	if err != nil{
		log.Fatalf(err.Error())
	}
	t.Run()
}
