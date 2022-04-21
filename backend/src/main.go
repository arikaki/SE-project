package main

import (
	"log"
	"net/http"

	// "encoding/json"

	// "github.com/dgrijalva/jwt-go"

	// "github.com/gorilla/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"kora.com/project/src/controllers"
	"kora.com/project/src/database"
)

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Establishing Database Connection
	database.DBConnect()
	// Route handles & endpoints

	r.HandleFunc("/login", database.Login).Methods("POST")
	r.HandleFunc("/api/user/insert", database.InsertUsers).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("GET")

	apiRouter := r.PathPrefix("/api").Subrouter()
	controllers.MainController(apiRouter)

	corsWrapper := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Origin", "Accept", "Authorization", "*"},
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:3000"},
	})

	// Start server
	log.Fatal(http.ListenAndServe(":8000", corsWrapper.Handler(r)))
}
