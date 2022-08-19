package home

import (
	"github.com/gin-gonic/gin"
	"github.com/hecterbonha/sergeant/lib/shared"
)

func Routes(route *gin.Engine) {
	home := route.Group("/")
	home.Use(shared.I18n())
	{
		home.GET("/ping", PingHandler)
	}
}
