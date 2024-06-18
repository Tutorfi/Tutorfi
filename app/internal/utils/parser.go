package utils

import "github.com/labstack/echo/v4"


func PathParser(c echo.Context) string {
    path := c.Request().URL.Path
    return path
}

