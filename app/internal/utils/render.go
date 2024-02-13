package utils

import (
	"app/internal/public/components"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RenderComponents(ctx echo.Context, status int, t templ.Component, err error) error {

	ctx.Response().Writer.WriteHeader(status)
	if err != nil {
		temp := components.ErrorComponentMsg(err.Error(), t)
		err = temp.Render(ctx.Request().Context(), ctx.Response().Writer)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "failed to render response template")
		}
	}
	err = t.Render(ctx.Request().Context(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response template")
	}

	return nil
}

func RenderPages(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)
	err := t.Render(ctx.Request().Context(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response template")
	}

	return nil
}
