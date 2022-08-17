package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Ping : Ping Router
func (HomeRouter) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"token": "token",
		"user":  "user",
	})
}
