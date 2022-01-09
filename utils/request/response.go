package request

import (
	"encoding/json"
	"log"
	"net/http"
)

func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err)
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	switch statusCode {
	case http.StatusInternalServerError:
		err = json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
	case http.StatusNotFound:
		err = json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
	default:
		err = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}

	if err != nil {
		log.Println(err)
	}
}
