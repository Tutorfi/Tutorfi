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
	"github.com/google/uuid"
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
		return c.String(http.StatusOK, "Account already exists")
	}
	//Check and hash the password
	password := c.FormValue("password")
	if utf8.RuneCountInString(password) < 8{
		fmt.Println("Password too short")
		return c.String(http.StatusOK, "Invalid Password")
	}
	//In the future we may need a restriction on passwords too long
	//Create a new account
	
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil{
		fmt.Println("password hasing failed")
		fmt.Println(err)
		return c.String(http.StatusOK, "Invalid Password")
	}
	
	err = handle.store.CreateAccount(c.FormValue("fname"), c.FormValue("lname"), email, string(hash))
	if err != nil{
		fmt.Println(err)
		return c.String(http.StatusOK, "Unknown Account creation error")
	}
	fmt.Println("account created successfully")
	c.Response().Header().Set("HX-Redirect", "/login")
	return c.String(http.StatusCreated, "Created account")
}
func createCookie(sessionid string) *http.Cookie{
	var cookie = new(http.Cookie)
	cookie.Name = "UUID"
	cookie.Value = sessionid
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
		return c.String(http.StatusOK, "Could not find account")
	}
	hash := []byte (account.Password)
	if bcrypt.CompareHashAndPassword(hash, []byte (password)) == nil{
		sessionid := uuid.New()
		cookie := createCookie(sessionid.String())
		c.SetCookie(cookie)
		handle.store.SetSessionID(email, sessionid.String())
		c.Response().Header().Set("HX-Redirect", "/")
		return c.String(http.StatusOK, "Logged in")
	}
	fmt.Println("login failed")
	return c.String(http.StatusOK, "Invalid username or password")
}