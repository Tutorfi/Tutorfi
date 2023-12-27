package main

import (
	"app/internal/account"
	"app/internal/connection"
	"app/internal/pages"
	"fmt"
	"os"

	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

// Used to initialize the server
// func buildServer() {

// }

// Adds routes to the echo server
func addRoutes(e *echo.Echo, db *sql.DB) {
	accountController := account.NewAccountController(account.NewAccountModel(db))
	account.AddAccountRoutes(e, accountController)
	pages.AddPagesRoutes(e)
}

// Add a function that checks for flags here
// func parseArgs() {

func main() {
	e := echo.New()
	fmt.Println("Current Working Directory:")
	fmt.Println(os.Getwd())

	// Change this to use the env file and this doesn't work
	db, err := connection.ConnectPgsql()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
	// defer db.Close()

	addRoutes(e, db)

	e.Use(middleware.CORS())
	// e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}
