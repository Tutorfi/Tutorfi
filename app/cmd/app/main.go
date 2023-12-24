package main

import (
	"app/internal/models"
	"app/internal/pages"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	fmt.Println("Current Working Directory:")
	fmt.Println(os.Getwd())

	// Change this to use the env file and this doesn't work
	models.BootstrapMongo("mongodb://localhost:27017", "test", 30)

	pages.AddPagesRoutes(e)
	e.Use(middleware.CORS())
	// e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8000"))
}
