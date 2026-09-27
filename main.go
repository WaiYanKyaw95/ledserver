package main

import (
	"ledserver/handlers"
	"ledserver/middleware"
	"ledserver/serial"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	db := initDB()
	defer db.Close()

	port, err := serial.OpenPort(9600, "COM7")
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
