package account

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	//"github.com/markbates/goth" For future oauth stuff
	//"github.com/markbates/goth/gothic" //Wrapper for goth, should make it easier to use
)

type AccountController struct {
	model *accountModel
}

/*
AddAcountroutes add sign in route
	Prob takes in a parameter for the database connection
*/
func (t *AccountController) signIn(c echo.Context) error {
	fmt.Println("Got a GET request")
	correct, err := t.model.checkAccount("user", "pass")
	if err != nil {
		fmt.Println("Error checking account")
		fmt.Println(err)
		return err
	}

	// This doesn't contain the sign in page
	// create the sign in page in the pages folder
	if correct {
		err := c.Redirect(http.StatusMovedPermanently, "<URL>")
		return err
	}

	htmlstr := "Incorrect Username or Password"
	tmpl, err := template.New("t").Parse(htmlstr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return tmpl.Execute(c.Response().Writer, nil)
}

func (t *AccountController) validateUser(c echo.Context) error{
	fmt.Println("Got a Post request sign in")
	email := c.FormValue("email")
	password := c.FormValue("password")
	//Replace with actual validation later
	if email == "asdf" && password == "asdfqwe"{
		cookie := new(http.Cookie)
		cookie.Name = "sessionID"
		cookie.Value = "1" //Replace with an actual id later
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
		return c.Redirect(http.StatusOK, "<url>")
	}
	return c.String(http.StatusUnauthorized, "login failed")
}
func (t *AccountController) createAccount(c echo.Context) error {

	// This is example code
	correct, err := t.model.createAccount(53)
	if correct {
		// Add a success here
		return nil
	}
	if err != nil {
		fmt.Println("Error creating account")
		fmt.Println(err)
		return err
	}

	//Fix this
	htmlstr := "Incorrect Username or Password"
	tmpl, err := template.New("t").Parse(htmlstr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return tmpl.Execute(c.Response().Writer, nil)
}