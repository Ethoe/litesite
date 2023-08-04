package handlers

import (
	"cmd/server/main.go/internal/db"
	"cmd/server/main.go/pkg/api"
	"cmd/server/main.go/pkg/entities/users"
	"cmd/server/main.go/pkg/middleware"
	"net/http"
	"time"

	"github.com/gorilla/context"
)

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "ethoe_session",
		Value:    "",
		Path:     "/",
		Domain:   "ethoe.dev",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Now().Add(-24 * time.Hour),
	}
	http.SetCookie(w, cookie)

	user, ok := context.Get(r, middleware.UserContextKey).(users.User)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		api.BodyMarshal(w, api.Response{"success": true}, http.StatusOK)
		return
	}

	if user.ID != 0 {
		err := users.DeleteUserSession(db.Master, user.ID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			api.BodyMarshal(w, api.Response{"success": false, "error": err}, http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	api.BodyMarshal(w, api.Response{"success": true}, http.StatusOK)
}
