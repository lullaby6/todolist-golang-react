package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"todolist-golang-react/src/api/models"
	"todolist-golang-react/src/api/routes"
	"todolist-golang-react/src/database"
)

func main() {
	database.Connect()
	if err := database.DB.AutoMigrate(&models.ToDo{}); err != nil {
		panic(fmt.Sprintf("[ERROR] Problem migrating database ToDo table (%s).", err))
	}

	if err := godotenv.Load(); err != nil {
		fmt.Printf("[ERROR] Problem loading .env file (%s).", err)
	}

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	routes.ToDoRoutes(api)

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public"))))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	PORT := ":" + os.Getenv("PORT")
	fmt.Println("Server on port " + os.Getenv("PORT"))

	log.Fatal(http.ListenAndServe(PORT, handler))
}
