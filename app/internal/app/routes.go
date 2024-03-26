package app

import (
	"app/internal/handlers/account"
	"app/internal/handlers/interface"
	"app/internal/handlers/scheduler"

	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo, a *App) {

	accountFunctions := accounthandler.New(a.store)
	schedulerFunctions := schedulerhandler.New(a.store)
    interfaceFunctions := interfacehandler.New(a.store)

	e.POST("/login/verify", accountFunctions.Verification)
	e.GET("/schedule/date", schedulerFunctions.Schedule)
	e.POST("/create-account/create", accountFunctions.CreateAccount)
    e.GET("/account/groups", interfaceFunctions.GetAccountGroups)

	e.Static("/css", "/app/internal/public/css")
	e.Static("/js", "/app/internal/public/js")
}
