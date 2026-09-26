package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

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
	var input LoginInput
	var user User

	// get the username and password
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check in the db if the username is there
	row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", input.Username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)

	// if not -> 401 invalid credentials
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
		// else database error 500
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// compare the password against stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	// if not matched -> 401 invalid credentials
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// generate a token for the sessions
	token := make([]byte, 32)
	_, err = rand.Read(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// store user id, token and created time in the sessions table
	// encode token into hex before storing it in the table
	encodedToken := hex.EncodeToString(token)
	_, err = db.Exec("INSERT INTO sessions (user_id, token, created_at) VALUES (?, ?, ?)", user.ID, string(encodedToken), time.Now().Format(time.RFC3339))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// response includes the session token, status code 200
	c.JSON(http.StatusOK, gin.H{"username": user.Username, "token": encodedToken})
}

func Logout(c *gin.Context, db *sql.DB) {
	var user User
	// get the session token from the authorization header
	bearer := c.GetHeader("Authorization")
	token := strings.TrimPrefix(bearer, "Bearer ")

	// check if the token is in the table
	row := db.QueryRow("SELECT user_id FROM sessions WHERE token = ?", token)
	// if not -> 401
	err := row.Scan(&user.ID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no valid token"})
		return
		// if nil, database error 500
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// delete the session token
	_, err = db.Exec("DELETE FROM sessions WHERE token = ?", token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// response 200
	c.JSON(http.StatusOK, gin.H{"success": "log out"})
}
