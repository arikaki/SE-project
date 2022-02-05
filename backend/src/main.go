package main

import (
	"log"
	"net/http"

	"kora.com/project/src/controllers"
	"kora.com/project/src/database"

	"github.com/gorilla/mux"
)

// Main function
func main() {
	// Init router
	r := mux.NewRouter()
	// Route handles & endpoints
	database.DBConnect()
	apiRouter := r.PathPrefix("/api").Subrouter()
	controllers.MainController(apiRouter)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
