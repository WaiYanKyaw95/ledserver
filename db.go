package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func initDB() *sql.DB {
	// create and connect to a database
	db, err := sql.Open("sqlite", "ledserver.db")
	if err != nil {
		log.Fatal(err)
	}

	// test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// create a user table if not exist yet
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	// create a session table if not exist yet
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INT NOT NULL,
		token TEXT NOT NULL UNIQUE,
		created_at TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	// return db
	return db
}
