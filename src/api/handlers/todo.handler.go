package handlers

import (
	"encoding/json"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	statusCode := 202

	statusText := http.StatusText(statusCode)
	w.WriteHeader(statusCode)

	data := map[string]interface{}{
		"status": statusText,
		"data":   "[]",
	}

	json.NewEncoder(w).Encode(data)
}
