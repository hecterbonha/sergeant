package routes

import (
	"github.com/hecterbonha/sergeant/routes/auth"
	"github.com/hecterbonha/sergeant/routes/home"
	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	home.HomeRouter{}.Init(g.Group(""))
	auth.AuthRouter{}.Init(g.Group("/auth"))
}
