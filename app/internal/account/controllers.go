package account

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
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
		c.Redirect(http.StatusMovedPermanently, "<URL>")
		return nil
	}

	htmlstr := fmt.Sprintf("Incorrect Username or Password")
	tmpl, err := template.New("t").Parse(htmlstr)

	return tmpl.Execute(c.Response().Writer, nil)
}
