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
}

func NewApp(listenAddr string, store storage.Storage) *App {
	return &App{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (a *App) Start(e *echo.Echo) error {
	pages.AddPagesRoutes(e)
	addRoutes(e, a)

	e.Use(middleware.CORS(), middleware.Logger(), middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 6}))
	return e.Start(a.listenAddr)
}
