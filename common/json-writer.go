package common

import (
	"encoding/json"
	"net/http"
)

func JsonWriter(w http.ResponseWriter, statusCode int, value any) error {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}
