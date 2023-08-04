package handlers

import (
	"cmd/server/main.go/internal/db"
	"cmd/server/main.go/pkg/api"
	"cmd/server/main.go/pkg/entities/files"
	"cmd/server/main.go/pkg/entities/users"
	"cmd/server/main.go/pkg/middleware"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

// TODO: Check total size user has taken up and cap that

/*
curl --location 'localhost:5050/api/file' \
--header 'Cookie: ethoe_session=y9feCnkE8HjTA7DOIAk0WdtNMzvAc10UtBHnbAsJX7w=' \
--form 'file=@"dAmeoGMGZ/trees.png"'
*/

const maxUploadSize = 10 << 30 // 1 GB max file size

func UploadFile(w http.ResponseWriter, r *http.Request) {
	user, ok := context.Get(r, middleware.UserContextKey).(users.User)
	if !ok {
		api.BodyMarshal(w, api.Response{"success": false, "error": "user not found"}, http.StatusUnauthorized)
		return
	}

	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		api.BodyMarshal(w, api.Response{"success": false, "error": err}, http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		api.BodyMarshal(w, api.Response{"success": false, "error": err}, http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileSize := fileHeader.Size
	if fileSize > maxUploadSize {
		api.BodyMarshal(w, api.Response{"success": false, "error": err}, http.StatusBadRequest)
		return
	}

	url, err := files.CreateFile(db.Master, file, fileHeader, user)
	if err != nil {
		log.Println(err)
		api.BodyMarshal(w, api.Response{"success": false, "error": err}, http.StatusInternalServerError)
		return
	}

	api.BodyMarshal(w, api.Response{"success": true, "url": url}, http.StatusOK)
}
