package middleware

import (
	"cmd/server/main.go/internal/db"
	"cmd/server/main.go/pkg/entities/users"

	"database/sql"
	"net/http"

	"github.com/gorilla/context"
)

type ContextKey string

const UserContextKey ContextKey = "user"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for the session token in the request cookie
		sessionToken, err := r.Cookie("ethoe_session")
		if err != nil || sessionToken.Value == "" {
			next.ServeHTTP(w, r)
			return
		}

		user, err := getUserBySessionToken(db.Master, sessionToken.Value)
		if err != nil {
			http.Error(w, "Invalid session token", http.StatusUnauthorized)
			return
		}

		context.Set(r, UserContextKey, user)

		next.ServeHTTP(w, r)
	})
}

func getUserBySessionToken(db *sql.DB, sessionToken string) (users.User, error) {
	// Query the database to retrieve user information using the session token
	var user users.User
	query := "SELECT u.id, u.username, u.email, u.password, u.reg_date " +
		"FROM usersessions s INNER JOIN users u ON s.userid = u.id WHERE s.sessiontoken = ?"
	err := db.QueryRow(query, sessionToken).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password, &user.RegDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return users.User{}, nil
		}
		return users.User{}, err
	}
	return user, nil
}
