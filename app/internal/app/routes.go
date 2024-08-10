package app

import (
	accountFunctions "app/internal/handlers/account"
	accounthandler "app/internal/handlers/account"
	"app/internal/handlers/app/account"
	"app/internal/handlers/app/interface"
	"app/internal/handlers/app/scheduler"
	filehandler "app/internal/handlers/file"
	interfacehandler "app/internal/handlers/interface"

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

func apiRoutes(e *echo.Echo, a *App) {
    interfaceFunctions := interfacehandler.New(a.store)
    accountFunctions := accounthandler.New(a.store)
    fileFunctions := filehandler.New(a.store, a.fileserver)

    e.GET("/api/group", interfaceFunctions.GetAccountGroups)
    e.POST("/api/account/create", accountFunctions.CreateAccount)
    e.POST("/api/account/login", accountFunctions.Login)
    e.GET("/api/account/verify", accountFunctions.VerifyCookie)
    e.GET("/api/account/logout", accountFunctions.Logout)

    e.GET("/api/account/group/get-file", fileFunctions.GetFile)
}
