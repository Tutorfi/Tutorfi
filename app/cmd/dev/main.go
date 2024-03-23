package main

import (
	"app/internal/app"
	storage "app/internal/storage/postgres"
	"app/internal/utils"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func proxy(e *echo.Echo) {
	url1, err := url.Parse("http://localhost:8000")
	if err != nil {
	e.Logger.Fatal(err)
	}	

	url2, err := url.Parse("http://localhost:8080")
	if err != nil {
	e.Logger.Fatal(err)
	}
	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
		{
			URL: url2,
		},
		}

	g := e.Group("/blog")
	g.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))
	g.GET("/thing", func(c echo.Context) error {		
		return c.String(http.StatusOK, "Hello");
	})
}

func main() {
	db, err := app.ConnectPgsql()
	if err != nil {
		fmt.Println(err)
	}
	postgresStorage := storage.NewPostgresStorage(db)

	build := flag.Bool("build", false, "ReBuild the database")
	flag.Parse()
	fmt.Println("Build:", *build)

	if *build {
		fmt.Println("Building database")
        err = postgresStorage.BuildDevDB()
        if err != nil {
            fmt.Println("Error building database")
            fmt.Println("Build failed")
        }
		os.Exit(0)
	}
	e := echo.New()
	fmt.Println("Current Working Directory:")
	fmt.Println(os.Getwd())
	if utils.GetEnv() == "prod" || utils.GetEnv() == "production" {
		fmt.Println("WARNING: Running in production mode")
	}

	proxy(e)
	
	// postgresStorage.BuildDevDB()
	server := app.NewApp("0.0.0.0:8000", postgresStorage)
	err = server.Start(e)

	e.Logger.Fatal(err)
}