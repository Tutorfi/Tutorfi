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
	fmt.Println(c.Cookies())
	if len(c.Cookies()) == 0{
		fmt.Println("no cookies")
		login := ""
		return tmpl.Execute(c.Response().Writer, login)
	}
	return tmpl.Execute(c.Response().Writer, nil)
	
}
func AddPagesRoutes(e *echo.Echo) {
	fmt.Println("Adding pages routes")

	e.GET("/", homePage)
	e.GET("/sign-in", signIn)
	e.GET("/login", login)
	e.GET("/create-account", createAccount)
}