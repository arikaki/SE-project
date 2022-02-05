package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is running"))
}

func MainController(r *mux.Router) {
	r.HandleFunc("/health-check", healthcheck).Methods("GET")
	UserController(r.PathPrefix("/user").Subrouter())
	QuestionController(r.PathPrefix("/question").Subrouter())
}
