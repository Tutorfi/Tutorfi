package pages

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func AddPagesRoutes(e *echo.Echo) {
	fmt.Println("Adding pages routes")

	e.GET("/", homePage)
    e.GET("/home", userHome)
	e.GET("/login", login)
	e.GET("/create-account", createAccount)
	e.GET("/test", test)
	e.GET("/about", aboutPage)
    e.GET("/group", groupView)
}
