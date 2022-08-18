package lib

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hecterbonha/sergeant/config"
	"github.com/hecterbonha/sergeant/lib/home"
)

func Server() {
	router := gin.Default()
	home.Routes(router)
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{config.RehearsalURL()},
		AllowMethods: []string{"GET", "POST"},
		MaxAge:       12 * time.Hour,
	}))
	router.Run(":8080")
}