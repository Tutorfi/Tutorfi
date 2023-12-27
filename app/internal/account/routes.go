package account

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// AddAcountroutes adds account option routes to the echo router
func AddAccountRoutes(e *echo.Echo, controller *AccountController) {
	fmt.Println("Adding account routes")

	e.GET("/sign-in", controller.signIn)
	e.GET("/sign-in", controller.createAccount)
}
