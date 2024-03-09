package interfacehandler

import (
	logintempl "app/internal/public/views/login"
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
	acc, err := utils.GetAccountFromSessionId(handle.store, c)
	if err == sql.ErrNoRows {
        return utils.RenderComponents(c, 200, logintempl.Error("Account Not Found"), nil)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(acc.Firstname)
	return utils.RenderComponents(c, 200, logintempl.Error("Account Created"), nil)
}
