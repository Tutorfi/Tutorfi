package pages

import (
	"app/internal/public/views/event"
	"app/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func event(c echo.Context) error {
	err := utils.RenderPages(c, http.StatusOK, eventtempl.Calendar([]string{"event/event.css"}))
	//Call the event handler
	if err != nil {
		return err
	}
	return nil

}
