package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
