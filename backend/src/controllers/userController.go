package controllers

import (
	"github.com/gorilla/mux"
	"kora.com/project/src/database"
)

func UserController(r *mux.Router) {
	r.HandleFunc("/insert", database.InsertUsers).Methods("POST")
	r.HandleFunc("/delete/{id}", database.InsertUsers).Methods("DELETE")
	r.HandleFunc("/update/{id}", database.InsertUsers).Methods("PUT")
	r.HandleFunc("/get/{id}", database.InsertUsers).Methods("GET")
}
