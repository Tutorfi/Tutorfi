package main

import (
	"app/internal/app"
	"app/internal/storage/postgres"
	"app/internal/utils"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	fmt.Println("Current Working Directory:")
	fmt.Println(os.Getwd())
	if utils.GetEnv() == "prod" || utils.GetEnv() == "production" {
		fmt.Println("WARNING: Running in production mode")
	}
	db, err := app.ConnectPgsql()

	if err != nil {
		fmt.Println(err)
	}

	server := app.NewApp("0.0.0.0:8000",storage.NewPostgresStorage(db)) 
	err = server.Start(e)

	e.Logger.Fatal(err)
}
