package controllers

import (
	"github.com/gorilla/mux"
	"kora.com/project/src/database"
)

func QuestionController(r *mux.Router) {
	r.HandleFunc("/ask", database.AskQ).Methods("POST")
	// r.HandleFunc("/edit", database.EditQ).Methods("POST")
	// r.HandleFunc("/get", database.GetQ).Methods("GET")
	r.HandleFunc("/search", database.FindMatchingQuestions).Methods("POST")
	r.HandleFunc("/getAll", database.GetAllQ).Methods("GET")
	r.HandleFunc("/getUnanswered", database.GetUnanswered).Methods("GET")
	r.HandleFunc("/topquestion", database.TopQuestion).Methods("POST")
	r.HandleFunc("/upvotequestion", database.UpvoteQuestion).Methods("GET")
	r.HandleFunc("/downvotequestion", database.DownvoteQuestion).Methods("GET")
	r.HandleFunc("/selectedquestion", database.SelectedQuestion).Methods("POST")
	r.HandleFunc("/report", database.Report).Methods("GET")
}
