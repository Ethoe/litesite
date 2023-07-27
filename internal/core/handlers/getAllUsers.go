package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	RegDate   string
}

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
