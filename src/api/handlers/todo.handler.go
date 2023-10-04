package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"todolist-golang-react/src/api/models"
	"todolist-golang-react/src/database"
)

func GetAllToDos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	statusCode := 200
	statusText := http.StatusText(statusCode)
	w.WriteHeader(statusCode)

	var todos []models.ToDo

	if err := database.DB.Find(&todos).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"status":   statusText,
		"endpoint": r.URL.Path,
		"data":     todos,
	}

	json.NewEncoder(w).Encode(data)
}

func CreateToDo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	statusCode := 200
	statusText := http.StatusText(statusCode)
	w.WriteHeader(statusCode)

	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var todo models.ToDo

	if title, ok := body["title"].(string); ok {
		todo.Title = title
	}

	if done, ok := body["done"].(bool); ok {
		todo.Done = done
	}

	if err := database.DB.Create(&todo).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"status":   statusText,
		"endpoint": r.URL.Path,
		"data":     body,
	}

	json.NewEncoder(w).Encode(data)
}

func UpdateToDo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	statusCode := 200
	statusText := http.StatusText(statusCode)
	w.WriteHeader(statusCode)

	vars := mux.Vars(r)

	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var todo models.ToDo

	if err := database.DB.First(&todo, vars["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if title, ok := body["title"].(string); ok {
		todo.Title = title
	}

	if done, ok := body["done"].(bool); ok {
		todo.Done = done
	}

	if err := database.DB.Save(&todo).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := database.DB.Model(models.ToDo{}).Where("id = ?", vars["id"]).Updates(&todo).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"status":   statusText,
		"endpoint": r.URL.Path,
		"ID":       vars["id"],
		"data":     body,
	}

	json.NewEncoder(w).Encode(data)
}

func DeleteToDo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	statusCode := 200
	statusText := http.StatusText(statusCode)
	w.WriteHeader(statusCode)

	vars := mux.Vars(r)

	if err := database.DB.Delete(&models.ToDo{}, "id = ?", vars["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"status":   statusText,
		"endpoint": r.URL.Path,
		"ID":       vars["id"],
	}

	json.NewEncoder(w).Encode(data)
}
