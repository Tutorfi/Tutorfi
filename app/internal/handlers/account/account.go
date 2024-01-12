/*
Contains the base handlers for creating and authenticating accounts.
*/
package accounthandler

import (
	"app/internal/storage"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"unicode/utf8"
	"time"
	"database/sql"
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
	//Get and check the email to see if the account exists
	email := c.FormValue("email")
	_, err := handle.store.GetAccount(email)
	if err != sql.ErrNoRows{
		fmt.Println("Account already exists")
		fmt.Println(err)
		return err
	}
	//Check and hash the password
	password := c.FormValue("password")
	if utf8.RuneCountInString(password) < 8{
		fmt.Println("Password too short")
		return c.String(http.StatusForbidden, "Invalid Password")
	}
	//In the future we may need a restriction on passwords too long
	//Create a new account
	
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil{
		fmt.Println("password hasing failed")
		fmt.Println(err)
		return nil
	}
	
	err = handle.store.CreateAccount(c.FormValue("fname"), c.FormValue("lname"), email, string(hash))
	if err != nil{
		fmt.Println(err)
		return c.String(http.StatusForbidden, "Account creation error")
	}
	fmt.Println("account created successfully")
	return c.Redirect(http.StatusFound, "/login")
}
func createCookie(altid string) *http.Cookie{
	var cookie = new(http.Cookie)
	cookie.Name = "UUID"
	cookie.Value = altid
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Path = "/"
	return cookie
}
func (handle *AccountHandler) Verification(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	fmt.Println("Login request")
	account, err := handle.store.GetAccount(email)
	if err != nil{
		fmt.Println(err)
		fmt.Println("No account found")
		return c.String(http.StatusForbidden, "Could not find account")
	}
	hash := []byte (account.Password)
	if bcrypt.CompareHashAndPassword(hash, []byte (password)) == nil{
		cookie := createCookie(account.ID)
		c.SetCookie(cookie)
		c.Response().Header().Set("HX-Redirect", "/")
		fmt.Println(c.Response().Header())
		return c.String(http.StatusOK, "Logged in")
	}
	fmt.Println("login failed")
	return c.String(http.StatusForbidden, "Invalid username or password")
}