package middleware

import (
	"database/sql"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int
	Username string
}

func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		// get token from the authorization header
		bearer := c.GetHeader("Authorization")
		token := strings.TrimPrefix(bearer, "Bearer ")

		// check if it's a valid token
		rowUserID := db.QueryRow("SELECT user_id FROM sessions WHERE token = ?", token)
		err := rowUserID.Scan(&user.ID)
		// if not -> 401 unauthorized, stop the request
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(401, gin.H{"error": "no valid token"})
			return
			// if it's server -> 500 database error
		} else if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "database error"})
			return
		}
		// if found -> attach the user to the context and continue
		// get the username from the users table
		rowUsername := db.QueryRow("SELECT username FROM users WHERE id = ?", user.ID)
		err = rowUsername.Scan(&user.Username)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "database error"})
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
