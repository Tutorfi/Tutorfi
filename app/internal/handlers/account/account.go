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
	"app/internal/public/views/login"
	"app/internal/utils"
	"github.com/go-playground/validator/v10"
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
	validate := validator.New(validator.WithRequiredStructEnabled())
	//Get and check the email to see if the account exists
	email := c.FormValue("email")
	err := validate.Var(email, "required,email")//Note that adding a space here will throw an error, maybe lint this???
	if err != nil{
		fmt.Println("Invalid email")
		fmt.Println(err)
		return utils.RenderComponents(c, 200, logintempl.Error(err), nil)
	}
	_, err = handle.store.GetAccount(email)
	if err != sql.ErrNoRows{
		fmt.Println("Account already exists")
		fmt.Println(err)
		return utils.RenderComponents(c, 200, logintempl.Error(err), nil)
	}

	//Check and hash the password
	password := c.FormValue("password")

	if utf8.RuneCountInString(password) < 8{
		fmt.Println("Password too short")
		return utils.RenderComponents(c, 200, logintempl.Error("Password must be longer than 8 characters"), nil)
	}
	//In the future we may need a restriction on passwords too long
	//Create a new account
	
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil{
		fmt.Println("password hasing failed")//Don't know when this will happen
		fmt.Println(err)
		return utils.RenderComponents(c, 200, logintempl.Error(err), nil)
	}
	
	err = handle.store.CreateAccount(c.FormValue("fname"), c.FormValue("lname"), email, string(hash))
	if err != nil{
		fmt.Println(err)
		return utils.RenderComponents(c, 200, logintempl.Error(err), nil)
	}
	err = validate.Struct(handle.store.GetAccount(email))
	if err != nil{
		fmt.Println(err)
		return utils.RenderComponents(c, 200, logintempl.Error(err), nil)
	}
	fmt.Println("account created successfully")
	c.Response().Header().Set("HX-Redirect", "/login")
	return utils.RenderComponents(c, 201, logintempl.Error("Account Created"), nil)
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
	validate := validator.New(validator.WithRequiredStructEnabled())
	email := c.FormValue("email")
	password := c.FormValue("password")
	
	fmt.Println("Login request")
	account, err := handle.store.GetAccount(email)
	if err != nil || validate.Struct(account) != nil{
		fmt.Println(err)
		fmt.Println(validate.Struct(account))
		fmt.Println("No account found")
		return utils.RenderComponents(c, 200, logintempl.Error("Invalid email or password"), nil)
	}
	hash := []byte (account.Password)
	if bcrypt.CompareHashAndPassword(hash, []byte (password)) == nil{
		//Create a new session id, set this session id in the database and make a cookie for it
		sessionid := uuid.New()
		cookie := createCookie(sessionid.String())
		c.SetCookie(cookie)
		err := handle.store.SetSessionID(email, sessionid.String())
		if err != nil{
			fmt.Println("cookie error")
			fmt.Println(err)
			return utils.RenderComponents(c, 200, logintempl.Error("Unknown error, please try again"), nil)
		}
		c.Response().Header().Set("HX-Redirect", "/")
		return utils.RenderComponents(c, 200, logintempl.Error("Logging in"), nil)
	}
	fmt.Println("login failed")
	return utils.RenderComponents(c, 200, logintempl.Error("Invalid email or password"), nil)
}