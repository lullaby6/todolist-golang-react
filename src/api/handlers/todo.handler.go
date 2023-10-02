package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllToDos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	statusCode := 200
	statusText := http.StatusText(statusCode)
	w.WriteHeader(statusCode)

	data := map[string]interface{}{
		"status": statusText,
		"data":   "[]",
	}

	json.NewEncoder(w).Encode(data)
}

func CreateToDo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	statusCode := 200
	statusText := http.StatusText(statusCode)
	w.WriteHeader(statusCode)

	data := map[string]interface{}{
		"status": statusText,
		"data":   "[]",
	}

	json.NewEncoder(w).Encode(data)
}

func DeleteToDo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	statusCode := 200
	statusText := http.StatusText(statusCode)
	w.WriteHeader(statusCode)

	vars := mux.Vars(r)

	data := map[string]interface{}{
		"status": statusText,
		"data":   "[]",
		"id":     vars["id"],
	}

	json.NewEncoder(w).Encode(data)
}
