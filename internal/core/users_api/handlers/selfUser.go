package handlers

import (
	"cmd/server/main.go/pkg/api"
	"cmd/server/main.go/pkg/entities/users"
	"cmd/server/main.go/pkg/middleware"
	"net/http"

	"github.com/gorilla/context"
)

type SelfUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func SelfUser(w http.ResponseWriter, r *http.Request) {
	user, ok := context.Get(r, middleware.UserContextKey).(users.User)
	if !ok {
		api.BodyMarshal(w, api.Response{"success": false, "error": "Cookie not found"}, http.StatusUnauthorized)
		return
	}

	if user.Email == "" {
		api.BodyMarshal(w, api.Response{"success": false, "error": "User not found"}, http.StatusUnauthorized)
		return
	}

	res := SelfUserResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	api.BodyMarshal(w, api.Response{"success": true, "user": res}, http.StatusOK)

}
