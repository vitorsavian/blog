package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	serv := echo.New()

	serv.Use()

	// Middleware
	serv.Use(middleware.Logger())
	serv.Use(middleware.Recover())

	// Routes
	Routes(serv)

	// Start server
	serv.Logger.Fatal(serv.Start(":80"))
}
