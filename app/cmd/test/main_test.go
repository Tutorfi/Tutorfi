package main

import (
	"app/internal/app"
	_ "github.com/jackc/pgx/v5"
	"testing"
)

func TestMain(t *testing.T) {
	t.Logf("works")
	_, err := app.BuildTestDB()
	if err != nil {
		t.Fatal(err)
	}
}
