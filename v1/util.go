package v1

import (
	"encoding/json"
	"net/http"
)

func renderJSON(w http.ResponseWriter, data interface{}, code int) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil
}
