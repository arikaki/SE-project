package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is running"))
}

func DBConnect() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func insertUsers(w http.ResponseWriter, r *http.Request) {
	client := DBConnect()
	type User struct {
		name     string
		email    string
		username string
		password string
		// following []
		// followers
		// topics
	}
	harshwardhan := User{"Harshwardhan", "harshwardhan0812@gmail.com", "SU", "password"}
	lokesh := User{"Lokesh", "loki@gmail.com", "Loki", "password"}
	nikhil := User{"Nikhil", "nik@gmail.com", "Nik", "password"}
	sairishab := User{"Sai Rishab", "SR@gmail.com", "SR", "password"}
	collection := client.Database("KoraDB").Collection("Users")
	users := []interface{}{harshwardhan, lokesh, nikhil, sairishab}

	insertManyResult, err := collection.InsertMany(context.TODO(), users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()
	// Route handles & endpoints

	r.HandleFunc("/health-check", getPosts).Methods("GET")
	r.HandleFunc("/insertUsers", insertUsers).Methods("GET")
	// r.HandleFunc("/books", getBooks).Methods("GET")
	// r.HandleFunc("/books/{id}", getBook).Methods("GET")
	// r.HandleFunc("/books", createBook).Methods("POST")
	// r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
