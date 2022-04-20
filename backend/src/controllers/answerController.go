package controllers

import (
	"github.com/gorilla/mux"
	"kora.com/project/src/database"
)

func AnswerController(r *mux.Router) {
	r.HandleFunc("/addAnswer", database.AddAnswer).Methods("POST")
}
