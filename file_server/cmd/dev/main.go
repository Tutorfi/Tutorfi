package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/:fileName", func(c echo.Context) error {		
		return c.File("/file_server/data/"+c.Param("fileName"))
	})
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}