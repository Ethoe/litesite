package handlers

import (
	"cmd/server/main.go/pkg/api"
	"cmd/server/main.go/pkg/entities/users"
	"net/http"
)

func SelfUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value("user").(users.User)
	if !ok {
		api.BodyMarshal(w, api.Response{"success": true, "error": "Cookie not found"}, http.StatusUnauthorized)
		return
	}

	api.BodyMarshal(w, api.Response{"success": true, "user": user}, http.StatusOK)

}
