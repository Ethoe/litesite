package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func AddUser(db *sql.DB, firstname, lastname, email, password string) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare("INSERT INTO users (firstname, lastname, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the statement to insert the new user
	_, err = stmt.Exec(firstname, lastname, email, password)
	if err != nil {
		return fmt.Errorf("failed to insert new user: %v", err)
	}

	return nil
}
