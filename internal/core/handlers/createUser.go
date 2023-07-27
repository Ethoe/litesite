package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cmd/server/main.go/internal/db"
	"cmd/server/main.go/pkg/entities/users"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user users.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	err = users.AddUser(db.Master, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
