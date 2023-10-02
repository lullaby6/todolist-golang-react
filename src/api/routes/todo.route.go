package routes

import (
	"github.com/gorilla/mux"

	"todolist-golang-react/src/api/handlers"
)

func ToDoRoutes(api *mux.Router) {
	r := api.PathPrefix("/todo").Subrouter()

	r.HandleFunc("", handlers.GetAllToDos).Methods("GET")
	r.HandleFunc("", handlers.CreateToDo).Methods("POST")
	r.HandleFunc("/{id}", handlers.DeleteToDo).Methods("DELETE")
}
