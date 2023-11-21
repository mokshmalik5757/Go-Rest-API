package server

import "github.com/gin-gonic/gin"

func HealthCheckRoute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "server is OK"})
	})
}