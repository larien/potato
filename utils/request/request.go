package request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func URLParam(r *http.Request, name string) string {
	return mux.Vars(r)[name]
}

func Decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("could not decode request body: %w", err)
	}
	return nil
}
