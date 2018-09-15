package webserver

import (
	"net/http"

	"github.com/labstack/echo"
)

// CreateWebServer configures an iris application
func CreateWebServer() *echo.Echo {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Lets test this")
	})

	return e
}
