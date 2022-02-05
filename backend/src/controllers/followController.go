package controllers

import (
	"github.com/gorilla/mux"
	"kora.com/project/src/database"
)

func FollowController(r *mux.Router) {
	r.HandleFunc("/", database.InsertUsers).Methods("POST")
	r.HandleFunc("/topic", database.InsertUsers).Methods("POST")
	r.HandleFunc("/user", database.InsertUsers).Methods("POST")
}
