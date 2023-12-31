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
	"bcrypt"
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
	email := c.FormValue("email")
	account, err := handle.store.GetAccount(email)
	if account{
		fmt.Println("Account already exists")
		return nil
	}
	password := c.FormValue("password")
	if RuneCountInString(password) < 8{
		fmt.Println("Password too short")
		return nil
	}
	if RuneCountInString(password) > 72{
		fmt.Println("Password too long")
		return nil
	}
	password := c.FormValue("password")
	hash, err := GenerateFromPassword(password, RuneCountInString(password))
	handle.store.CreateAccount(c.FormValue("fname"), c.FormValue("lname"), email, hash)
	return nil
}

func (handle *AccountHandler) Verification(c echo.Context) error {
	// Pull the data from the request
	email := c.FormValue("email")
	account, err := handle.store.GetAccount(email)
	if err != nil {
		fmt.Println("No account found")
		return err
	}

	fmt.Println(account)
	fmt.Println(handle.store)
	
	// Check if the account password matches the hashed password
	
	matched := false
	password := c.FormValue("password")
	hash := handle.store.GetPassword(account)
	matched, _ := CompareHashAndPassword(hash, password)
	if matched {
		err := c.Redirect(http.StatusOK, "<URL>")
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