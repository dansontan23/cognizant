package db

import (
	"database/sql"
	"log"
)

// Declare a global variable to hold the database connection
var DB *sql.DB

// Initialize SQLite database connection
func DbInit() *sql.DB {
	var err error
	DB, err = sql.Open("sqlite3", "../elibrary.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	log.Println("Successfully started db")
	return DB
}

// Close the database connection when the application ends
func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Fatal("Error closing database:", err)
		}
		log.Println("Successfully closed db connection")
	}
}
