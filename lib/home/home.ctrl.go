package home

import (
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	nameFromQuery := c.Query("name")

	message := GetPingMessages(nameFromQuery)

	c.JSON(200, gin.H{
		"message": message,
		"status":  "OK",
	})
}
