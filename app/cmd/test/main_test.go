package main

import (
	"app/internal/app"
	_ "github.com/jackc/pgx/v5"
	"testing"
)

func TestMain(t *testing.M) {
	app.BuildTestDB()
	t.Run()
}
