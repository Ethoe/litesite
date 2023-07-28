package handlers

import (
	"fmt"
	"io"
	"net/http"

	"cmd/server/main.go/internal/db"
	"cmd/server/main.go/pkg/entities/users"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := users.GetAllUsers(db.Master)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get users: %v", err), http.StatusBadRequest)
		return
	}
	io.WriteString(w, fmt.Sprintf("%v", users))
}
