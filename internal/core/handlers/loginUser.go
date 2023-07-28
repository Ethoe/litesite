package handlers

import (
	"cmd/server/main.go/internal/db"
	"cmd/server/main.go/pkg/entities/users"
	"encoding/json"
	"net/http"
	"time"
)

type loginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the username and password
	var login loginData
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := users.SignInUser(db.Master, login.Email, login.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    token,
		Expires:  time.Now().Add(365 * 24 * time.Hour), // Set the expiration time for the session cookie
		HttpOnly: true,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"success": "true"})
}