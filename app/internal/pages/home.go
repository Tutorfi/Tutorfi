package pages

import (
	"fmt"
	"html/template"
	"github.com/labstack/echo/v4"
)

func homePage (c echo.Context) error {
	fmt.Println("Got a GET request")
	temp, e := template.ParseFiles("./public/index.html")
	if e != nil {
		fmt.Println("Error parsing template")
		fmt.Println(e)
		return e
	}
	tmpl := template.Must(temp, e)
	return tmpl.Execute(c.Response().Writer, nil)
}

func AddPagesRoutes(e *echo.Echo) {
	fmt.Println("Adding pages routes")

	e.GET("/", homePage)

}