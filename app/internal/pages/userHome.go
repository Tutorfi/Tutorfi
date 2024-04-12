package pages

import (
	userinterface "app/internal/public/views/interface"
	"app/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

//  All logic is handled in the handlers package ex: geting users data
func userHome(c echo.Context) error {
    err := utils.RenderPages(c, http.StatusOK, userinterface.ViewOrganizations())
    if err != nil {
        return err
    }
    return nil
}
