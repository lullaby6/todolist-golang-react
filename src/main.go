package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"todolist-golang-react/src/api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file.")
	}

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	routes.ToDoRoutes(api)

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public"))))

	PORT := ":" + os.Getenv("PORT")
	fmt.Println("Server on port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(PORT, r))
}
