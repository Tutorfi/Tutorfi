package pages

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
    fmt.Print(code)
	c.Logger().Error(err)
    c.Logger().Print("hello1");
	fmt.Sprintf("hello")
    // errorPage := 
	// if err := c.File(errorPage); err != nil {
	// 	c.Logger().Error(err)
	// }
}


func AddPagesRoutes(e *echo.Echo, user *echo.Group) {
    fmt.Println("Adding pages routes")
    e.GET("/", homePage)
    e.GET("/login", login)
    e.GET("/create-account", createAccount)
    e.GET("/test", test)
    e.GET("/about", aboutPage)
    e.GET("/:errors", errorHandler)
    e.HTTPErrorHandler = customHTTPErrorHandler

    user.GET("/test", test)
}
