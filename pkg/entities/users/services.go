package users

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers(db *sql.DB) ([]User, error) {
	// Prepare the SQL query
	rows, err := db.Query("SELECT id, firstname, lastname, email, password, reg_date FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}
	defer rows.Close()

	var users []User

	// Iterate over the result set and add each user to the slice
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.RegDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration: %v", err)
	}

	return users, nil
}

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
