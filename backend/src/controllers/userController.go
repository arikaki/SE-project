package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"kora.com/project/src/database"
)

func UserController(r *mux.Router) {
	r.HandleFunc("/delete/{id}", database.InsertUsers).Methods("DELETE")
	r.HandleFunc("/update/{id}", database.InsertUsers).Methods("PUT")
	r.HandleFunc("/fetch-user", database.FetchUser).Methods("POST")
	r.HandleFunc("/dummyanswer", database.InsertDummyAnswer).Methods("GET")
	r.HandleFunc("/dummyuser", database.InsertDummyUser).Methods("GET")
	r.HandleFunc("/dummyquestion", database.InsertDummyQuestion).Methods("GET")
	r.HandleFunc("/deleteuser", database.DeleteUser).Methods("GET")
	r.HandleFunc("/topquestion", database.TopQuestion).Methods("POST")
	r.HandleFunc("/upvoteanswer", database.UpvoteAnswer).Methods("GET")
	r.HandleFunc("/downvoteanswer", database.DownvoteAnswer).Methods("GET")
	r.HandleFunc("/upvotequestion", database.UpvoteQuestion).Methods("GET")
	r.HandleFunc("/downvotequestion", database.DownvoteQuestion).Methods("GET")
}

// func InsertDummyData(w http.ResponseWriter, r *http.Request) {
// 	// database.InsertDummyUser()
// 	database.InsertDummyQuestion()
// }

func Logout(w http.ResponseWriter, r *http.Request) {

	expiration := time.Now().Add(-24 * time.Hour)
	cookie := http.Cookie{Name: "Login", Value: "", Expires: expiration}
	http.SetCookie(w, &cookie)

	jsonResponse, err := json.Marshal("Logout Successful")
	if err != nil {
		return
	}
	//update response
	w.Write(jsonResponse)
}
