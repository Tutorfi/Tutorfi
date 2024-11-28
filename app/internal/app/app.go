package app

import (
	"app/internal/pages"
	"app/internal/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	listenAddr string
	store      storage.Storage
    fileserver string
}

func NewApp(listenAddr string, store storage.Storage, fileserver string) *App {
	return &App{
		listenAddr: listenAddr,
		store:      store,
        fileserver: fileserver,
	}
}

func (a *App) AuthMiddleware (next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cookie, err := c.Cookie("Tutorfi_Account")
        if err != nil {
            return c.Redirect(302, "/login")
        }
        sessionId := cookie.Value
        acc, err := a.store.GetAccountSessionId(sessionId)
        if err != nil {
            // Oops something happened
            return c.Redirect(302, "/login")
        }
        if acc == nil {
            // Unauthorized
            return c.Redirect(302, "/login")
        }
        return next(c)
    }
}

func (a *App) Start(e *echo.Echo) error {
    user := e.Group("/user", a.AuthMiddleware)
    pages.AddPagesRoutes(e, user)
	addRoutes(e, a)
    apiRoutes(e, a)

	e.Use(middleware.CORS(), middleware.Logger(), middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 6}))
	return e.Start(a.listenAddr)
}

