package lib

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hecterbonha/sergeant/config"
	"github.com/hecterbonha/sergeant/db"
	"github.com/hecterbonha/sergeant/lib/auth"
	"github.com/hecterbonha/sergeant/lib/home"
)

func Server() {
	gin.SetMode(config.GinMode())
	db.ConnectionInit()
	app := gin.Default()
	app.SetTrustedProxies(nil)
	home.Routes(app)
	auth.Routes(app)
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{config.RehearsalURL()},
		AllowMethods: []string{"GET", "POST"},
		MaxAge:       12 * time.Hour,
	}))
	app.Run(config.SergeantURL())
}
