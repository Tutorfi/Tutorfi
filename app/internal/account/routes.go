package account

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// AddAcountroutes adds account option routes to the echo router
func AddAcountroutes(e *echo.Echo, controller *AccountController) {
	fmt.Println("Adding account routes")

	e.POST("/sign-in", controller.validateUser)
	e.GET("/sign-in", controller.signIn)
	e.GET("/sign-in", controller.createAccount)
}
