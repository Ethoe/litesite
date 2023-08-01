package handlers

import (
	"net/http"
	"time"
)

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		Domain:   "ethoe.dev",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Now().Add(-24 * time.Hour),
	}
	http.SetCookie(w, cookie)
}
