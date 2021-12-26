package handler

import "net/http"

func V1GetPotato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		_, _ = w.Write([]byte("{\"content\": \"potato\"}"))
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	_, _ = w.Write([]byte("{\"error\": \"method not allowed\"}"))
}
