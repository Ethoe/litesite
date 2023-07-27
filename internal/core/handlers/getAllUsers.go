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
	users, _ := users.GetAllUsers(db.Master)
	io.WriteString(w, fmt.Sprintf("%v", users))
}
