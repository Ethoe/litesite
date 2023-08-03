package handlers

import (
	"cmd/server/main.go/internal/db"
	"cmd/server/main.go/pkg/api"
	"cmd/server/main.go/pkg/entities/files"
	"cmd/server/main.go/pkg/entities/users"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value("user").(users.User)
	if !ok {
		api.BodyMarshal(w, api.Response{"success": false, "error": "user not found"}, http.StatusUnauthorized)
		return
	}

	// Get the file ID from the URL parameter
	vars := mux.Vars(r)
	fileIDstr := vars["id"]
	fileID, err := strconv.Atoi(fileIDstr)
	if err != nil {
		api.BodyMarshal(w, api.Response{"success": false, "error": "invalid file id"}, http.StatusBadRequest)
		return
	}

	err = files.DeleteFile(db.Master, fileID, user)
	if err != nil {
		log.Println(err)
		api.BodyMarshal(w, api.Response{"success": false, "error": err}, http.StatusInternalServerError)
		return
	}

	api.BodyMarshal(w, api.Response{"success": true}, http.StatusOK)
}
