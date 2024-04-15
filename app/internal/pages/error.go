package pages

import (
	status "app/internal/public/views/status"
	"app/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func errorHandler(c echo.Context) error {
	err := utils.RenderPages(c, http.StatusOK, status.ErrorPage())
	if err != nil {
		return err
	}
	return nil
}
