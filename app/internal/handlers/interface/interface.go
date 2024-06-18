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
        c.Response().Header().Add("HX-Redirect", "/login")
	    return c.Redirect(302,"/login")
    }
	if err != nil {
		fmt.Println(err)
<<<<<<< HEAD
        re := fillResponse("Failed", "None", "Server error, try again later", nil, "")
	    return c.JSON(http.StatusInternalServerError, re)
=======
        c.Response().Header().Add("HX-Redirect", "/login")
	    return c.Redirect(302,"/login")
>>>>>>> 3bbfde8 (added_basic_queries)
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
