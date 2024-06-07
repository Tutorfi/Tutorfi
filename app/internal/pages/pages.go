package pages

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddPagesRoutes(e *echo.Echo, user *echo.Group) {
	fmt.Println("Adding pages routes")
	e.GET("/", homePage)
	e.GET("/home", userHome)
	e.GET("/login", login)
	e.GET("/create-account", createAccount)
	e.GET("/test", test)
	e.GET("/about", aboutPage)
	e.GET("/:errors", errorHandler)
	e.GET("/group", groupView)
	e.GET("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, c.Request().URL.Path)
	})
	user.GET("/test", test)
}
