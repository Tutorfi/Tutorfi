package pages

import (
	userinterface "app/internal/public/views/interface"
	"app/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func groupView(c echo.Context) error {
    err := utils.RenderPages(c,http.StatusOK, userinterface.ViewGroups())
    if err != nil {
        return nil
    }
    return nil
}
