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
		t.Errorf(`Expected a string of user details. Got '%v'`, resp.Body.String())
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
	t.Errorf(`Expected the product name "User is Deleted". Got '%v'`, resp.Body.String())

}

func Test_FindMatchingQuestions(t *testing.T) {
	var jsonStr = []byte(`{"Search": "why my investment doesn't work"}`)
	req, _ := http.NewRequest("POST", "/search", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	a := http.HandlerFunc(database.FindMatchingQuestions)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	fmt.Println("resp", resp.Body.String())

	if resp.Body.String() != `[[{"Key":"_id","Value":"62217301ecf350ef0c2e0dc6"},{"Key":"question","Value":"What’s a good investment for 2022?"}],[{"Key":"_id","Value":"62475a4bca2bb7bc2c1d960b"},{"Key":"question","Value":"What’s a good investment for long term?"}],[{"Key":"_id","Value":"62475ac0ca2bb7bc2c1d960d"},{"Key":"question","Value":"What’s a good investment to double the money in 14 days?"}]]` {
		t.Errorf(`Expected list of matching questions. Got '%v'`, resp.Body.String())
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
		t.Errorf(`Expected a string of questions. Got '%v'`, resp.Body.String())
	}
}

func Test_GetUnanswered(t *testing.T) {
	req, _ := http.NewRequest("GET", "/getUnanswered", nil)
	req.Header.Set("Content-Type", "application/json")

	a := http.HandlerFunc(database.GetUnanswered)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	fmt.Println("resp", resp.Body.String())

	if resp.Body.String() != `[[{"Key":"_id","Value":"61fdd6999ff44333f800c14d"},{"Key":"question","Value":"Why me?"},{"Key":"user","Value":"61fdd6999ff44333f800c14c"},{"Key":"upvotes","Value":23},{"Key":"downvotes","Value":1},{"Key":"is_answered","Value":false},{"Key":"followers","Value":[]},{"Key":"topics","Value":[]},{"Key":"answers","Value":[]},{"Key":"comments","Value":[]}],[{"Key":"_id","Value":"61fdf8c7cc1711bfb99ae5f8"},{"Key":"question","Value":"Why Software Engineering?"},{"Key":"user","Value":"61fdf8c7cc1711bfb99ae5f7"},{"Key":"upvotes","Value":100},{"Key":"downvotes","Value":56},{"Key":"is_answered","Value":false},{"Key":"followers","Value":[]},{"Key":"topics","Value":[]},{"Key":"answers","Value":[]},{"Key":"comments","Value":[]}]]` {
		t.Errorf(`Expected a list of all unanswered questions and Got '%v'`, resp.Body.String())
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
		t.Errorf(`Expected a string of answers. Got '%v'`, resp.Body.String())
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
