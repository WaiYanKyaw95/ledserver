package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context, db *sql.DB) {
	var input LoginInput
	// get the username and password
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash the password using bcrypt
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	// insert username and hashed password into the users table
	// password needs to be converted into a string
	_, err := db.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`, input.Username, string(hash))
	if err != nil {
		// if unique constraint error -> 409 username already taken
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			c.JSON(http.StatusConflict, gin.H{"error": "username already taken"})
			return
		}
		// else database error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// response 201 created
	c.JSON(http.StatusCreated, gin.H{"username": input.Username})
}

func Login(c *gin.Context, db *sql.DB) {
	// get the username and password
	// check in the db if the username is there
	// if not -> 401 invalid credentials
	// compare the password against stored hashed password
	// if not matched -> 401 invalid credentials
	// generate a token for the sessions
	// store user id, token and created time in the sessions table
	// response includes the session token, status code 200
}

func Logout(c *gin.Context, db *sql.DB) {
	// get the session token from the request
	// check if the token is in the table
	// if not -> 404 user not logged in
	// delete the session token
	// response 200
}
