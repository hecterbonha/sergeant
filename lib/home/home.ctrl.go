package home

import (
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	nameFromQuery := c.Query("name")
	if nameFromQuery == "" {
		c.JSON(400, gin.H{
			"message": "MALFORMED_REQUEST_FORMAT",
			"status":  "NOT_OK",
		})
		return
	}

	message := GetPingMessages(nameFromQuery)

	c.JSON(200, gin.H{
		"message": message,
		"status":  "OK",
	})
}
