/*
Contains the base handlers for creating and authenticating accounts.
*/
package accounthandler

import (
	"app/internal/storage"
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	store storage.Storage
}

func New(store storage.Storage) *AccountHandler {
	return &AccountHandler{
		store: store,
	}
}

func (handle *AccountHandler) CreateAccount(c echo.Context) error {
	return nil
}

func (handle *AccountHandler) Verification(c echo.Context) error {
	// Pull the data from the request
	account, err := handle.store.GetAccount("thing@gmail.com")
	if err != nil {
		return err
	}

	fmt.Println(account)
	
	// Check if the account password matches the hashed password
	
	matched := false
	if matched {
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