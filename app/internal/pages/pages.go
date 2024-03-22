package pages

import (
	"app/internal/storage"
	"fmt"

	"github.com/labstack/echo/v4"
)

func AddPagesRoutes(e *echo.Echo, s storage.Storage) {
    fmt.Println("Adding pages routes")

    e.GET("/", homePage)
    e.GET("/login", login)
    e.GET("/create-account", createAccount)
    e.GET("/test", test)
    e.GET("/about", aboutPage)

    user := e.Group("/user", s.authMiddleware)
}
