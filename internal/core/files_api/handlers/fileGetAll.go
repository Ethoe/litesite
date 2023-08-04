package handlers

import (
	"cmd/server/main.go/internal/db"
	"cmd/server/main.go/pkg/api"
	"cmd/server/main.go/pkg/entities/files"
	"cmd/server/main.go/pkg/entities/users"
	"cmd/server/main.go/pkg/middleware"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
)

func GetAllFiles(w http.ResponseWriter, r *http.Request) {
	user, ok := context.Get(r, middleware.UserContextKey).(users.User)
	if !ok {
		api.BodyMarshal(w, api.Response{"success": false, "error": "user not found"}, http.StatusUnauthorized)
		return
	}

	// Parse the query parameters for pagination
	limitStr := r.URL.Query().Get("limit")
	pageStr := r.URL.Query().Get("page")

	// Convert the limit and page parameters to integers
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20 // Default limit if not specified or invalid
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 0 {
		page = 0 // Default page if not specified or invalid
	}

	// Get all files associated with the user
	fileList, err := files.GetFilesByUserID(db.Master, user, limit, page)
	if err != nil {
		log.Println(err)
		api.BodyMarshal(w, api.Response{"success": false, "error": err}, http.StatusInternalServerError)
		return
	}

	api.BodyMarshal(w, api.Response{"success": true, "files": fileList}, http.StatusOK)
}
