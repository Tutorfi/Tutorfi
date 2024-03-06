package pages

import (
	abouttempl "app/internal/public/views/about"
	"app/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func aboutPage(c echo.Context) error {
	err := utils.RenderPages(c, http.StatusOK, abouttempl.About())
	if err != nil {
		return err
	}
	return nil
}
