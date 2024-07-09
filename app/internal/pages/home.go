package pages

import (
	"app/internal/public/components"
	"app/internal/public/views/home"
	"app/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

//  All logic is handled in the handlers package ex: geting users data
func homePage(c echo.Context) error {
	err := utils.RenderPages(c, http.StatusOK, hometempl.Home())
	if err != nil {
		return err
	}
	return nil
}

func test(c echo.Context) error {
	return utils.RenderPages(c, http.StatusOK, components.Testing("Test"))
}

