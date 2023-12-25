package account

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Implement the db template here

/*
AddAcountroutes add sign in route
	Prob takes in a parameter for the database connection
*/
func signIn(c echo.Context) error {
	fmt.Println("Got a GET request sign in")
	temp, e := template.ParseFiles("./public/sign-in.html")
	if e != nil {
		fmt.Println("Error parsing template")
		fmt.Println(e)
		return e
	}
	tmpl := template.Must(temp, e)
	return tmpl.Execute(c.Response().Writer, nil)
}

func validateUser(c echo.Context) error{
	fmt.Println("Got a Post request sign in")
	email := c.FormValue("email")
	password := c.FormValue("password")
	//Replace with actual validation later
	//tmpl := template.Must(template.ParseFiles("./public/sign-in.html"))
	
	if email == "asdf" && password == "asdfqwe"{
		//tmpl.ExecuteTemplate(c.Response().Writer, "login-error", "login failed")
		fmt.Println("success")
		return c.Redirect(http.StatusOK, "/pages/index.html")
		
	}
	return c.String(http.StatusUnauthorized, "login failed")
}