package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"kora.com/project/src/auth"
	"kora.com/project/src/database"
)

type login struct {
	Email    string `json:"Email"`
	Username string `json:"UserName"`
	Password string `json:"Password"`
}

var SecretKey = []byte(os.Getenv("Cookie_Key"))

func UserController(r *mux.Router) {
	r.HandleFunc("/delete/{id}", database.InsertUsers).Methods("DELETE")
	r.HandleFunc("/update/{id}", database.InsertUsers).Methods("PUT")
	r.HandleFunc("/get-all-details", database.FetchUsers).Methods("GET")
	//r.HandleFunc("/fetch-user", database.FetchUsers).Methods("GET")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data login

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fetchedUser, err := database.GetUser(data.Username)

	if err := bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(data.Password)); err != nil {
		//err
	}
	fmt.Println("fetchedUser", fetchedUser.Username)
	claims := &auth.Claims{
		Username: fetchedUser.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SecretKey)
	fmt.Println("token", tokenString, SecretKey)

	if err != nil {
		// err
	}

	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{Name: "Session", Value: tokenString, Expires: expiration, HttpOnly: true}
	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	//convert struct to JSON
	jsonResponse, err := json.Marshal("Login Successful")
	if err != nil {
		return
	}
	//update response
	w.Write(jsonResponse)
}

func Logout(w http.ResponseWriter, r *http.Request) {

	expiration := time.Now().Add(-24 * time.Hour)
	cookie := http.Cookie{Name: "Login", Value: "", Expires: expiration, HttpOnly: true}
	http.SetCookie(w, &cookie)

	jsonResponse, err := json.Marshal("Logout Successful")
	if err != nil {
		return
	}
	//update response
	w.Write(jsonResponse)
}
