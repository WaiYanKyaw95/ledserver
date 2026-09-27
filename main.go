package main

import (
	"ledserver/handlers"
	"ledserver/middleware"
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
		// default to /dev/arduino
		// during development in Windows, set export SERIAL_PORT=COM7 before go run .
		portName = "/dev/arduino"
	}

	router.StaticFile("/", "./static/index.html")

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
			handlers.LedOn(c, portName)
		})
		protected.POST("/led/off", func(c *gin.Context) {
			handlers.LedOff(c, portName)
		})
	}

	router.Run(":8080")
}
