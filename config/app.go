package config

import "github.com/gin-gonic/gin"

func RehearsalURL() string {
	i := Config("REHEARSAL_URL")

	return i
}

func GinMode() string {
	g := Config("GIN_MODE")
	if g == "release" {
		return gin.ReleaseMode
	}
	return gin.DebugMode
}
