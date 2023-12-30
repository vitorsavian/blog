package server

import (
	"github.com/labstack/echo/v4"
	"github.com/vitorsavian/dandd/bahamut/pkg/html"
	view "github.com/vitorsavian/dandd/bahamut/web/index"
)

func index(c echo.Context) error {
	return html.HTML(c, view.Hello("World"))
}
