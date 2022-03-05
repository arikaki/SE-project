package controllers

import (
	"fmt"
	"testing"

	"bytes"
	"net/http"
	"net/http/httptest"
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
