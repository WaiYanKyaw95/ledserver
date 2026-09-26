package handlers

import "github.com/gin-gonic/gin"

func LedOn(c *gin.Context) {
	c.JSON(200, gin.H{"message": "LED on"})
}

func LedOff(c *gin.Context) {
	c.JSON(200, gin.H{"message": "LED off"})
}
