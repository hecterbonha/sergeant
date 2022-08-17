package lib

import (
	"github.com/hecterbonha/sergeant/lib/auth"
	"github.com/hecterbonha/sergeant/lib/home"
	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	home.HomeRouter{}.Init(g.Group(""))
	auth.AuthRouter{}.Init(g.Group("/auth"))
}
