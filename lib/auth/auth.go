package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hecterbonha/sergeant/lib/shared"
)

func Routes(route *gin.Engine) {
	auth := route.Group("/auth")
	auth.Use(shared.I18n())
	{
		auth.GET("/generate", Generate)
		auth.GET("/handshake", HandShake)
	}
}
