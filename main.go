package main

import (
	"ledserver/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	db := initDB()
	defer db.Close()

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

	router.Run(":8080")
}
