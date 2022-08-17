package main

import (
	"github.com/hecterbonha/sergeant/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.Routes(e.Group(""))

	e.Logger.Fatal(e.Start(":8080"))
}
