package config

import (
	"github.com/gin-gonic/gin"
)

func SergeantURL() string {
	s := Config("SERGEANT_URL")

	return s
}

func RehearsalURL() string {
	s := Config("REHEARSAL_URL")

	return s
}

func GinMode() string {
	g := Config("GIN_MODE")
	if g == "release" {
		return gin.ReleaseMode
	}
	return gin.DebugMode
}

func PostgresURL() string {
	p := Config("POSTGRES_URL")
	return p
}
