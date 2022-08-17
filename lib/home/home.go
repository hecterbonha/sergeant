package home

import (
	"github.com/labstack/echo/v4"
)

// HomeRouter : HomeRouter struct
type HomeRouter struct{}

// Init : Init Router
func (ctrl HomeRouter) Init(g *echo.Group) {
	g.GET("/ping", ctrl.Ping)
}
