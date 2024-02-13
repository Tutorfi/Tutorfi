/*
Contains the base handlers for creating and authenticating accounts.
*/
package accounthandler

import (
	"app/internal/public/views/createAccount"
	"app/internal/public/views/login"
	"app/internal/storage"
	"app/internal/utils"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"unicode/utf8"
	"time"
	"database/sql"
	"github.com/google/uuid"
	"app/internal/public/views/login"
	"app/internal/utils"
	"regexp"
	"net/mail"
)

type AccountHandler struct {
	store storage.Storage
}

func New(store storage.Storage) *AccountHandler {
	return &AccountHandler{
		store: store,
	}
}

func checkFormValue(expression, val string) error {
	regex := expression
	matcher, err := regexp.Compile(regex)
	if err != nil{
		fmt.Println("regex failed")//What to do here?
		return err
	}
	matched := matcher.MatchString(val)
	if !matched{
		return &AccountError{msg: fmt.Sprintf("Invalid form value")}
	}
	return nil
}
func (handle *AccountHandler) CreateAccount(c echo.Context) error {
	//Get and check the email to see if the account exists
	email := c.FormValue("email")
	_, err := mail.ParseAddress(email)
	if err != nil{
		fmt.Println("Invalid email")
		fmt.Println(email)
		return utils.RenderComponents(c, 200, logintempl.Error(err.Error()), nil)
	}
	_, err = handle.store.GetAccount(email)
	if err != sql.ErrNoRows{
		fmt.Println("Account already exists")
		fmt.Println(err)
        // Future: Change this to show server error and on dev show server error
		return utils.RenderComponents(c, 200, createAccountTempl.CreateAccountForm(form, "Invalid email or password", true), nil)
	}

	nameRegex := `^[A-Za-z\x{00C0}-\x{00FF}][A-Za-z\x{00C0}-\x{00FF}\'\-]+([\ A-Za-z\x{00C0}-\x{00FF}][A-Za-z\x{00C0}-\x{00FF}\'\-]+)*`
	fname := c.FormValue("fname")
	lname := c.FormValue("lname")
	err = checkFormValue(nameRegex, fname)
	if err != nil{
		return utils.RenderComponents(c, 200, logintempl.Error(err.Error()), nil)
	}
	err = checkFormValue(nameRegex, lname)
	if err != nil{
		return utils.RenderComponents(c, 200, logintempl.Error(err.Error()), nil)
	}

	//Check and hash the password
	//For the future when we figure out error handeling better
	//https://stackoverflow.com/questions/19605150/regex-for-password-must-contain-at-least-eight-characters-at-least-one-number-a

	if utf8.RuneCountInString(password) < 8{
		fmt.Println("Password too short")
		err := &AccountError{msg: fmt.Sprintf("Password must be longer than 8 characters")}
		return utils.RenderComponents(c, 200, logintempl.Error(err.Error()), nil)
	}
	//In the future we may need a restriction on passwords too long
	//Create a new account

	hash, err := bcrypt.GenerateFromPassword([]byte(form.Password), 0)
	if err != nil {
		fmt.Println("password hasing failed")
        fmt.Println(err)
		return utils.RenderComponents(c, 200, createAccountTempl.CreateAccountForm(form, "Invalid email or password", true), nil)
	}

	err = handle.store.CreateAccount(form.Fname, form.Lname, form.Email, string(hash))
	if err != nil {
		fmt.Println(err)
		return utils.RenderComponents(c, 200, logintempl.Error(err.Error()), nil)
	}
	//For now we allow all kinds of names, this may change in the future
	//https://stackoverflow.com/questions/2385701/regular-expression-for-first-and-last-name
	//https://andrewwoods.net/blog/2018/name-validation-regex/

	fmt.Println("account created successfully")
	c.Response().Header().Set("HX-Redirect", "/login")
	return utils.RenderComponents(c, 201, logintempl.Error("Account Created"), nil)
}
func createCookie(sessionid string) *http.Cookie {
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
	err := checkFormValue(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, email)
	if err != nil{
		return utils.RenderComponents(c, 200, logintempl.Error(err.Error()), nil)
	}

	account, err := handle.store.GetAccount(email)

	if err != sql.ErrNoRows {
		return utils.RenderComponents(c, 200, logintempl.Login(email, "Invalid email or password", true), nil)
	}

	if err != nil {
		fmt.Println(err)
	}
	hash := []byte(account.Password)

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
        return utils.RenderComponents(c, 200, logintempl.Login(email, "Invalid email or password", true), nil)
	}

	//Create a new session id, set this session id in the database and make a cookie for it
	sessionid := uuid.New()
	cookie := createCookie(sessionid.String())
	c.SetCookie(cookie)
	err = handle.store.SetSessionID(email, sessionid.String())
	if err != nil { //What to do here?
		fmt.Println("cookie error")
		fmt.Println(err)
		return utils.RenderComponents(c, 200, logintempl.Error(err.Error()), nil)
	}
	c.Response().Header().Set("HX-Redirect", "/")
    return utils.RenderComponents(c, 200, logintempl.Login("", "Logging in", false), nil)
}
