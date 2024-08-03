package interfacehandler

import (
	"app/internal/storage"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type InterfaceHandler struct {
	store storage.Storage
}

func New(store storage.Storage) *InterfaceHandler {
	return &InterfaceHandler{
		store: store,
	}
}

func (handle *InterfaceHandler) GetAccountGroups(c echo.Context) error {
    cookie, err := c.Cookie("Tutorfi_Account")
    if err != nil {
        return c.Redirect(302, "/login")
    }
    sessionId := cookie.Value
    acc, err := handle.store.GetAccountSessionId(sessionId)
	if err == sql.ErrNoRows {
	    return c.Redirect(302,"/login")
    }
	if err != nil {
		fmt.Println(err)
        re := fillResponse("Failed", "None", "Server error, try again later", nil, "")
	    return c.JSON(http.StatusInternalServerError, re)
	}
    groups, err := handle.store.GetGroups(acc)
    if err != nil {
        fmt.Println(err)
        re := fillResponse("Failed", "None", "Server error, try again later", nil, "")
	    return c.JSON(http.StatusInternalServerError, re)
    }
    if len(groups) == 0 {
        // Future update the "None" to the auth level of the tags
        re := fillResponse("Success", "None", "No groups", nil, "")
	    return c.JSON(http.StatusOK, re)
    }
    re := fillResponse("Success", "None", "", groups, "name")
	return c.JSON(http.StatusOK, re)
    
}

func (handle *InterfaceHandler) CreateGroup(c echo.Context) error {
    cookie, err := c.Cookie("Tutorfi_Account")
    u := new(response)
    err = c.Bind(u)
    if err != nil {
        fmt.Println(err)
        re := fillResponse("Invalid", "None", "No cookies", nil, "")
        return c.JSON(http.StatusUnauthorized, re)
    }

    sessionId := cookie.Value
    acc, err := handle.store.GetAccountSessionId(sessionId)
    
	if err == sql.ErrNoRows {
	    return c.Redirect(302,"/login")
    }
	if err != nil {
		fmt.Println(err)
    re := fillResponse("Failed", "None", "Server error, try again later", nil, "")
	    return c.JSON(http.StatusInternalServerError, re)
	}

    re := fillResponse("Success", "None", "created group", nil, "name")
    return c.JSON(http.StatusOK, re);
}
