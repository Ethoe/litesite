package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"cmd/server/main.go/internal/db"
	"cmd/server/main.go/pkg/api"
	"cmd/server/main.go/pkg/entities/users"
)

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user UserRegister
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	token, err := users.AddUser(db.Master, user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed save user to db: %v", err), http.StatusBadRequest)
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
	api.BodyMarshal(w, api.Response{"success": true, "user": user}, http.StatusCreated)
}
