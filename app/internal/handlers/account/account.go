/*
Contains the base handlers for creating and authenticating accounts.
*/
package accounthandler

import (
	"app/internal/public/views/createAccount"
	"app/internal/storage"
	"database/sql"
	"fmt"
	"unicode/utf8"

	"net/http"
	"net/mail"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	form := createAccountTempl.AccountForm{}
    u := new(accountAuth)
    err := c.Bind(u)

    if err != nil {
        re := fillResponse("Failed", "Server error, try again later")
        fmt.Println(err)
        return c.JSON(http.StatusInternalServerError, re)
    }

	form.Email = u.Email
	form.Fname = u.Fname
	form.Lname = u.Lname
	form.Password = u.Password

    _, err = mail.ParseAddress(form.Email)
	if err != nil {
        re := fillResponse("Invalid", "Invalid email")
        return c.JSON(http.StatusUnprocessableEntity, re)
	}
	_, err = handle.store.GetAccount(form.Email)
	if err != sql.ErrNoRows {
		// Future: Change this to show server error and on dev show server error
        re := fillResponse("Invalid", "Account already exists")
        return c.JSON(http.StatusConflict, re)
	}

	nameRegex := `^[A-Za-z\x{00C0}-\x{00FF}][A-Za-z\x{00C0}-\x{00FF}\'\-]+([\ A-Za-z\x{00C0}-\x{00FF}][A-Za-z\x{00C0}-\x{00FF}\'\-]+)*`
	err = checkFormValue(nameRegex, form.Fname)
	if err != nil {
		// Future: Change this to show server error and on dev show server error
        re := fillResponse("Invalid", "Invalid characters in first name")
        return c.JSON(http.StatusUnprocessableEntity, re)
	}
	err = checkFormValue(nameRegex, form.Lname)
	if err != nil {
        re := fillResponse("Invalid", "Invalid characters in last name")
        return c.JSON(http.StatusUnprocessableEntity, re)
	}

	//Check and hash the password
	//For the future when we figure out error handeling better
	//https://stackoverflow.com/questions/19605150/regex-for-password-must-contain-at-least-eight-characters-at-least-one-number-a
	if utf8.RuneCountInString(form.Password) < 8 {
        re := fillResponse("Invalid", "Invalid number of characters")
        return c.JSON(http.StatusUnprocessableEntity, re)
	}
	//In the future we may need a restriction on passwords too long
	//Create a new account
	hash, err := bcrypt.GenerateFromPassword([]byte(form.Password), 0)
	if err != nil {
		fmt.Println("password hashing failed")
		fmt.Println(err)
        re := fillResponse("Failed", "Server error, try again later")
        return c.JSON(http.StatusInternalServerError, re)
	}

	err = handle.store.CreateAccount(form.Fname, form.Lname, form.Email, string(hash))
	if err != nil {
		fmt.Println(err)
        re := fillResponse("Failed", "Server error, try again later")
        return c.JSON(http.StatusInternalServerError, re)
	}
	//For now we allow all kinds of names, this may change in the future
	//https://stackoverflow.com/questions/2385701/regular-expression-for-first-and-last-name
	//https://andrewwoods.net/blog/2018/name-validation-regex/
    re := fillResponse("Success", "Account Created")
    return c.JSON(http.StatusOK, re)
}


func (handle *AccountHandler) Verification(c echo.Context) error {
    u := new(accountAuth)
    err := c.Bind(u)
    if err != nil {
        re := fillResponse("Failed", "Server error, try again later")
        fmt.Println(err)
        return c.JSON(http.StatusInternalServerError, re)
    }

	email := u.Email
    password := u.Password

	err = checkFormValue(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, email)
	if err != nil {
		fmt.Println("email error")
        re := fillResponse("Invalid", "Invalid email or password")
        return c.JSON(http.StatusForbidden, re)
	}

	account, err := handle.store.GetAccount(email)

	if err == sql.ErrNoRows {
		fmt.Println(err)
        re := fillResponse("Invalid", "Invalid email or password")
        return c.JSON(http.StatusForbidden, re)
	}

	if err != nil {
		fmt.Println(err)
        re := fillResponse("Failed", "Server error, try again later")
        return c.JSON(http.StatusInternalServerError, re)
	}
	hash := []byte(account.Password)

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		fmt.Println(err)
        re := fillResponse("Failed", "Server error, try again later")
        return c.JSON(http.StatusInternalServerError, re)
	}

	//Create a new session id, set this session id in the database and make a cookie for it
	sessionid := uuid.New()
	cookie := createCookie(sessionid.String())
	c.SetCookie(cookie)
	err = handle.store.SetSessionID(email, sessionid.String())
	if err != nil { //What to do here?
		fmt.Println("cookie error")
		fmt.Println(err)
        // return a json error
	}
    re := fillResponse("Success", "Authorized")
    return c.JSON(http.StatusOK, re)
}
