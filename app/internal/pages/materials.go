package pages

import (
	materialstempl "app/internal/public/views/materials"
	"app/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func materialsPage (c echo.Context) error {
    err := utils.RenderPages(c, http.StatusOK, materialstempl.Materials())
	if err != nil {
		return err
	}
	return nil
}