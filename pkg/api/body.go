package api

import (
	"encoding/json"
	"net/http"
)

type Response map[string]interface{}

func BodyMarshal(w http.ResponseWriter, x map[string]interface{}, status int) {
	resp, err := json.Marshal(x)
	if err != nil {
		w.Write([]byte(`{"success":false,"message":"Some internal error occurred"}`))
	}

	w.WriteHeader(status)
	w.Write(resp)
}
