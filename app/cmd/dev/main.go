package main

import (
	"app/internal/app"
	"app/internal/storage/postgres"
	"app/internal/utils"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"golang.org/x/tools/go/analysis/passes/defers"
)

func main() {
	db, err := storage.ConnectPgsql()
	if err != nil {
		fmt.Println(err)
	}
    defer db.Close()
	postgresStorage := storage.NewPostgresStorage(db)

	build := flag.Bool("build", false, "Rebuild the database")
	flag.Parse()
	fmt.Println("Build:", *build)

	if *build {
		fmt.Println("Building database")
		err = postgresStorage.BuildDevDB()
		if err != nil {
			fmt.Println("Error building database")
			fmt.Println(err)
			fmt.Println("Build failed")
			os.Exit(1)
		}
		os.Exit(0)
	}

	e := echo.New()
	fmt.Println("Current Working Directory:")
	fmt.Println(os.Getwd())
	if utils.GetEnv() == "prod" || utils.GetEnv() == "production" {
		fmt.Println("WARNING: Running in production mode")
	}

	// postgresStorage.BuildDevDB()
	server := app.NewApp("0.0.0.0:8000", postgresStorage, "0.0.0.0:8080")
	err = server.Start(e)

	e.Logger.Fatal(err)
}
