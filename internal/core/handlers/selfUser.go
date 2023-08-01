package handlers

import (
	"cmd/server/main.go/pkg/api"
	"cmd/server/main.go/pkg/entities/users"
	"net/http"
)

type SelfUserResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func SelfUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value("user").(users.User)
	if !ok {
		api.BodyMarshal(w, api.Response{"success": true, "error": "Cookie not found"}, http.StatusUnauthorized)
		return
	}

	res := SelfUserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	api.BodyMarshal(w, api.Response{"success": true, "user": res}, http.StatusOK)

}
