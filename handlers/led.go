package handlers

import (
	serialpkg "ledserver/serial"

	"github.com/gin-gonic/gin"
)

func controlLED(c *gin.Context, portName string, command string, message string) {
	port, err := serialpkg.OpenPort(9600, portName)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not connect to Arduino"})
		return
	}
	defer port.Close()

	// send command to the port
	err = serialpkg.SendCommand(port, command)
	// handle error
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to send command"})
		return
	}
	c.JSON(200, gin.H{"message": message})
}

func LedOn(c *gin.Context, portName string) {
	controlLED(c, portName, "ON", "LED on")
}

func LedOff(c *gin.Context, portName string) {
	controlLED(c, portName, "OFF", "LED off")
}
