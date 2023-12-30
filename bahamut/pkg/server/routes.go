package server

import (
	"github.com/labstack/echo/v4"
)

func Routes(mux *echo.Echo) {
	// View routes
	mux.GET("/", index)

	// API routes
}
