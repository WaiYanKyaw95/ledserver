package handlers

import (
	serialpkg "ledserver/serial"

	"github.com/gin-gonic/gin"
	"go.bug.st/serial"
)

func LedOn(c *gin.Context, port serial.Port) {
	// send command to the port "ON"
	err := serialpkg.SendCommand(port, "ON")
	// handle error
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to send command"})
		return
	}
	c.JSON(200, gin.H{"message": "LED on"})
}

func LedOff(c *gin.Context, port serial.Port) {
	// send command to the port "OFF"
	err := serialpkg.SendCommand(port, "OFF")
	// handle error
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to send command"})
		return
	}
	c.JSON(200, gin.H{"message": "LED off"})
}
