package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Master *sql.DB

func SetupDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDB := os.Getenv("MYSQL_DB")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", mysqlUser, mysqlPassword, mysqlHost, mysqlDB)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	Master = db
}
