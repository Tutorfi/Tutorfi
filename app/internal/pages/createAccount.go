package pages

import (
	"app/internal/public/views/createAccount"
	"app/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func createAccount(c echo.Context) error {
	cssDir := []string{"account/account.css"}
	err := utils.RenderPages(c, http.StatusOK, createtempl.Create(cssDir))
	if err != nil {
		return err
	}
	return nil
}
