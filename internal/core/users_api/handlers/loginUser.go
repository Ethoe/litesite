package handlers

import (
	"cmd/server/main.go/internal/db"
	"cmd/server/main.go/pkg/api"
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

	token, user, err := users.SignInUser(db.Master, login.Email, login.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "ethoe_session",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(365 * 24 * time.Hour), // Set the expiration time for the session cookie
		HttpOnly: true,
	})

	w.Header().Set("Content-Type", "application/json")
	api.BodyMarshal(w, api.Response{"success": true, "user": user}, http.StatusOK)
}
