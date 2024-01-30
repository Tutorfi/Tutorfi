package pages

import (
	"app/internal/public/views/calendar"
	"app/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func calendar (c echo.Context) error {
	err := utils.RenderPages(c, http.StatusOK, calendartempl.Calendar([]string{"account/account.css"}))
	if err != nil {
		return err
	}
	return nil
}