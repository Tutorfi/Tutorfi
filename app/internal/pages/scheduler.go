package pages

import (
	schedulertempl "app/internal/public/views/scheduler"
	"app/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func schedulerPage(c echo.Context) error {
	err := utils.RenderPages(c, http.StatusOK, schedulertempl.Scheduler())
	if err != nil {
		return err
	}
	return nil
}
