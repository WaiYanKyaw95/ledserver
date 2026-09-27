package main

import (
	"ledserver/handlers"
	"ledserver/middleware"
	"ledserver/serial"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	db := initDB()
	defer db.Close()

	portName := os.Getenv("SERIAL_PORT")
	if portName == "" {
		// default for pi
		// during development in Windows, set export SERIAL_PORT=COM7 before go run .
		portName = "/dev/ttyACM0"
	}

	port, err := serial.OpenPort(9600, portName)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ledserver is running"})
	})

	router.POST("/register", func(c *gin.Context) {
		handlers.Register(c, db)
	})

	router.POST("/login", func(c *gin.Context) {
		handlers.Login(c, db)
	})

	router.POST("/logout", func(c *gin.Context) {
		handlers.Logout(c, db)
	})

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware(db))
	{
		protected.POST("/led/on", func(c *gin.Context) {
			handlers.LedOn(c, port)
		})
		protected.POST("/led/off", func(c *gin.Context) {
			handlers.LedOff(c, port)
		})
	}

	router.Run(":8080")
}
