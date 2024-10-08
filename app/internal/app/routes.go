package app
import (
	accounthandler "app/internal/handlers/account"
	"app/internal/handlers/app/account"
	"app/internal/handlers/app/interface"
	"app/internal/handlers/app/scheduler"
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
    accounthandler := accounthandler.New(a.store)

    e.GET("/api/group", interfaceFunctions.GetAccountGroups)
    e.POST("/api/account/create", accounthandler.CreateAccount)
    e.POST("/api/account/login", accounthandler.Login)
    e.GET("/api/account/verify", accounthandler.VerifyCookie)
    e.GET("/api/account/logout", accounthandler.Logout)
}
