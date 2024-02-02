package app

import (
	"app/internal/pages"
	"app/internal/storage"

	"github.com/labstack/echo/v4"
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
    addRoutes(e,a)

	return e.Start(a.listenAddr)
}
