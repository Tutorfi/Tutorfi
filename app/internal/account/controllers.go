package account

import (
	"fmt"
	"html/template"

	"github.com/labstack/echo/v4"
)

// Implement the db template here

/*
AddAcountroutes add sign in route
	Prob takes in a parameter for the database connection
*/
func signIn(c echo.Context) error {
	fmt.Println("Got a GET request")
	temp, e := template.ParseFiles("./public/sign-in.html")
	if e != nil {
		fmt.Println("Error parsing template")
		fmt.Println(e)
		return e
	}
	tmpl := template.Must(temp, e)
	return tmpl.Execute(c.Response().Writer, nil)
}
