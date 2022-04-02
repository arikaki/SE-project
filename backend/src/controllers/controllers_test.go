package controllers

import (
	"fmt"
	"testing"

	"bytes"
	"net/http"
	"net/http/httptest"

	"kora.com/project/src/database"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func Test_Login(t *testing.T) {
	var jsonStr = []byte(`{"Email": "harshwardhan0812@gmail.com", "Username": "SU", "Password": "password"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	a := http.HandlerFunc(Login)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	fmt.Println("resp", resp.Body.String())

	if resp.Body.String() != `"Login Successful"` {
		t.Errorf(`Expected product name to be "Login Successful". Got '%v'`, resp.Body.String())
	}
}

func Test_FetchUser(t *testing.T) {
	var jsonStr = []byte(`{"Username": "SR"}`)
	req, _ := http.NewRequest("POST", "/fetch-user", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	a := http.HandlerFunc(database.FetchUser)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	fmt.Println("resp", resp.Body.String())

	if resp.Body.String() != `{"Fullname":"SaiRishab","Email":"SR@gmail.com","UserName":"SR","Password":"","Topics":[],"Upvotes":0,"Downvotes":0,"Questions":null,"Answer":null}` {
		t.Errorf(`Expected product name to be "Login Successful". Got '%v'`, resp.Body.String())
	}
}

func Test_DeleteUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "/deleteuser", nil)
	req.Header.Set("Content-Type", "application/json")

	a := http.HandlerFunc(database.DeleteUser)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	fmt.Println("resp", resp.Body.String())

	if resp.Body.String() != `"User is deleted"` {
		t.Errorf(`Expected response to be "User is deleted". Got '%v'`, resp.Body.String())
	}
}
func Test_TopQuestion(t *testing.T) {
	var jsonStr = []byte(`{"topic": "Technology"}`)
	req, _ := http.NewRequest("POST", "/topquestion", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	a := http.HandlerFunc(database.DeleteUser)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	fmt.Println("resp", resp.Body.String())

	if resp.Body.String() != `{"question":"How can modern technology help evolve?","question":"What does “absolute refractive index of glass is 1.5” mean?","question":"what is the speed of the bullet train?"}` {
		t.Errorf(`Expected product name to be "Login Successful". Got '%v'`, resp.Body.String())
	}
}
func Test_SelectedQuestion(t *testing.T) {
	var jsonStr = []byte(`{"_id": "62217301ecf350ef0c2e0dc5"}`)
	req, _ := http.NewRequest("POST", "/selectedquestion", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	a := http.HandlerFunc(database.DeleteUser)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	fmt.Println("resp", resp.Body.String())

	if resp.Body.String() != `{"answer":"Technology has helped us to fly, drive, sail,communicate"}` {
		t.Errorf(`Expected product name to be "Feteched Sucessfully". Got '%v'`, resp.Body.String())
	}
}

func Test_Logout(t *testing.T) {
	req, _ := http.NewRequest("GET", "/logout", nil)
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(Logout)
	handler.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	if resp.Body.String() != `"Logout Successful"` {
		t.Errorf("handler returned unexpected body : got %v want %v", resp.Body.String(), `"Logout Successful"`)
	}
}
