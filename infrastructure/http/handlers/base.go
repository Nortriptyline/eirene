package handlers

import "github.com/gin-gonic/gin"

func CreatePingRoute(router *gin.RouterGroup) {
	// Implementation for creating a ping route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
