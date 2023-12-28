package main

import (
	"app/internal/app"
	"app/internal/storage/postgres"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

// Used to initialize the server
// func buildServer() {

// }

// Adds routes to the echo server
// func addRoutes(e *echo.Echo, db *sql.DB) {
// 	accountController := account.NewAccountController(account.NewAccountModel(db))
// 	account.AddAccountRoutes(e, accountController)
// 	pages.AddPagesRoutes(e)
// }

// Add a function that checks for flags here
// func parseArgs() {

func main() {
	e := echo.New()
	fmt.Println("Current Working Directory:")
	fmt.Println(os.Getwd())
	db, err := app.ConnectPgsql()

	if err != nil {
		fmt.Println(err)
	}

	test := app.NewApp("0.0.0.0:8000",storage.NewPostgresStorage(db)) 
	err = test.Start(e)

	// addRoutes(e, db)

	// e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}
