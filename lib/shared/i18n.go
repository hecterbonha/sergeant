package shared

import (
	"github.com/gin-gonic/gin"
)

func I18n() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := c.Request.Header["Accept-Language"]
		c.Set("locales", l[0])
		c.Next()
	}
}
