package interfacehandler

import (
	"app/internal/public/views/interface"
	"app/internal/storage"
	"app/internal/utils"
	"database/sql"
	"fmt"

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
	    return c.Redirect(302,"/login")
	}
    groups, err := handle.store.GetGroups(acc)
    if err != nil {
        fmt.Println(err)
        return utils.RenderComponents(c, 200, userinterface.NoGroups(),nil)
    }
    if len(groups) == 0 {
        return utils.RenderComponents(c, 200, userinterface.NoGroups(),nil)
    }
    return utils.RenderComponents(c, 200, userinterface.GroupComponent(groups),nil)
    
}
