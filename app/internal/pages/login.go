package pages

import (
	"app/internal/public/views/login"
	"app/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func login(c echo.Context) error {
	err := utils.RenderPages(c, http.StatusOK, logintempl.Login([]string{"account/account.css"}))
	if err != nil {
		return err
	}
	return nil
}
