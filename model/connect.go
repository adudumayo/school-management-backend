package model

import (
	"log"
	"os"
	"database/sql"
)

var db *sql.DB

// Connect to MySQL database
func connectDB() *sql.DB {
	var err error
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbName := os.Getenv("DBNAME")
	dbHost := os.Getenv("DBHOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	fmt.Println("Connected to the database!")

	return db
}

