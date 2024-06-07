package app

import (
	"app/internal/handlers/app/account"
	"app/internal/handlers/app/interface"
	"app/internal/handlers/app/scheduler"

	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo, a *App) {

	accountFunctions := appAccounthandler.New(a.store)
	schedulerFunctions := appschedulerhandler.New(a.store)
    interfaceFunctions := appinterfacehandler.New(a.store)

	e.POST("/login/verify", accountFunctions.Verification)
	e.GET("/schedule/date", schedulerFunctions.Schedule)
	e.POST("/create-account/create", accountFunctions.CreateAccount)
    e.GET("/account/groups", interfaceFunctions.GetAccountGroups)

	e.Static("/css", "/app/internal/public/css")
	e.Static("/js", "/app/internal/public/js")
}
