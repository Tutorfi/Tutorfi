package pages

import (
	"app/internal/public/components"
	hometempl "app/internal/public/views/home"
	"app/internal/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func homePage(c echo.Context) error {
	err := utils.RenderPages(c, http.StatusOK, hometempl.Home())
	if err != nil {
		return err
	}
	// Check if user is logged in
	if len(c.Cookies()) == 0 {
		fmt.Println(c.Cookies())
		fmt.Println("No cookies found")
	}
	return nil
}

func test(c echo.Context) error {
	return utils.RenderPages(c, http.StatusOK, components.Testing("Test"))
}
func AddPagesRoutes(e *echo.Echo) {
	fmt.Println("Adding pages routes")

	e.GET("/", homePage)
	e.GET("/login", login)
	e.GET("/create-account", createAccount)
	e.GET("/test", test)
	e.GET("/calendar", calendar)
	e.GET("/about", aboutPage)
}
