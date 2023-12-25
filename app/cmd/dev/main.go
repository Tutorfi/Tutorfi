package main

import (
	// "app/internal/models"
	"app/internal/pages"
	"fmt"
	"os"
	"app/internal/account"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Used to initialize the server
// func buildServer() {

// }

// Adds routes to the echo server
func addRoutes(e *echo.Echo) {
	pages.AddPagesRoutes(e)
	account.AddAcountroutes(e)
}

// Add a function that checks for flags here
// func parseArgs() {


func main() {
	e := echo.New()
	fmt.Println("Current Working Directory:")
	fmt.Println(os.Getwd())

	// Change this to use the env file and this doesn't work
	// db, err := models.BootstrapMongo(30)
	fmt.Println("Adding routes: ")
	addRoutes(e)

	e.Use(middleware.CORS())
	// e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8000"))
}
