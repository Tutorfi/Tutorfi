package main

import (
	"testing"
	_ "github.com/jackc/pgx/v5"
	"app/internal/app"
)

func TestMain(t *testing.T) {
	t.Logf("works")
	_, err := app.BuildTestDB()
	if err != nil{
		t.Fatalf(err.Error())
	}
}

