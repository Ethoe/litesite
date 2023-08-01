package users

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"

	"cmd/server/main.go/pkg/authentication"

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

func AddUser(db *sql.DB, firstname, lastname, email, password string) (string, error) {
	if firstname == "" || email == "" || password == "" {
		return "", fmt.Errorf("missing required fields")
	}

	if email != "ryanouttrim@gmail.com" {
		return "", fmt.Errorf("invalid email address")
	}

	hashedPassword := authentication.Hash(password)

	// Prepare the SQL statement
	stmt, err := db.Prepare("INSERT INTO users (firstname, lastname, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return "", fmt.Errorf("failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the statement to insert the new user
	_, err = stmt.Exec(firstname, lastname, email, hashedPassword)
	if err != nil {
		return "", fmt.Errorf("failed to insert new user: %v", err)
	}

	token, _, err := SignInUser(db, email, password)
	if err != nil {
		return "", fmt.Errorf("failed to sign in user: %v", err)
	}

	return token, nil
}

func SignInUser(db *sql.DB, email, password string) (string, User, error) {
	hashedPassword := authentication.Hash(password)
	var user User
	query := "SELECT id, firstname, lastname, email, password, reg_date " +
		"FROM users WHERE email = ? AND password = ?"
	err := db.QueryRow(query, email, hashedPassword).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.RegDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", User{}, fmt.Errorf("invalid email or password")
		}
		return "", User{}, err
	}

	token := generateSessionToken()
	query = "INSERT INTO usersessions (userid, sessiontoken, creation) VALUES (?, ?, NOW())"
	_, err = db.Exec(query, user.ID, token)
	if err != nil {
		return "", User{}, err
	}

	return token, user, nil
}

func generateSessionToken() string {
	// Generate a random session token using crypto/rand and base64 encoding
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		log.Fatal("Failed to generate random bytes:", err)
	}
	return base64.URLEncoding.EncodeToString(tokenBytes)
}
