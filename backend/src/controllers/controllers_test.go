package controllers

import (
	"context"
	"testing"

	"bytes"
	"net/http"
	"net/http/httptest"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

	a := http.HandlerFunc(database.Login)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

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

	if resp.Body.String() != ("User is Deleted") {
		t.Errorf(`Expected the product name "User is Deleted". Got '%v'`, resp.Body.String())
	}

}

func Test_FindMatchingQuestions(t *testing.T) {
	var jsonStr = []byte(`{"Search": "why my investment doesn't work"}`)
	req, _ := http.NewRequest("POST", "/search", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	a := http.HandlerFunc(database.FindMatchingQuestions)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	if resp.Body.String() != `[[{"Key":"_id","Value":"62217301ecf350ef0c2e0dc6"},{"Key":"question","Value":"What’s a good investment for 2022?"}],[{"Key":"_id","Value":"62475a4bca2bb7bc2c1d960b"},{"Key":"question","Value":"What’s a good investment for long term?"}],[{"Key":"_id","Value":"62475ac0ca2bb7bc2c1d960d"},{"Key":"question","Value":"What’s a good investment to double the money in 14 days?"}]]` {
		t.Errorf(`Expected list of matching questions. Got '%v'`, resp.Body.String())
	}
}
func Test_TopQuestion(t *testing.T) {
	var jsonStr = []byte(`{"topic": ["Chemistry"]}`)
	req, _ := http.NewRequest("POST", "/topquestion", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	a := http.HandlerFunc(database.TopQuestion)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	if resp.Body.String() != `[{"Key":"question","Value":"Is milk an organic compound?"},{"Key":"topic","Value":"Chemistry"}]` {
		t.Errorf(`Expected a string of questions. Got '%v'`, resp.Body.String())
	}
}
func Test_DownvoteQuestion(t *testing.T) {
	var jsonStr = []byte(`{"question":"Is milk an organic compound?"}`)
	req, _ := http.NewRequest("GET", "/downvotequestion", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	a := http.HandlerFunc(database.DownvoteQuestion)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	if resp.Body.String() != `"Downvote Question Successful"` {
		t.Errorf(`Expected message "Downvote Question Successful". Got '%v'`, resp.Body.String())
	}
}
func Test_UpvoteQuestion(t *testing.T) {
	var jsonStr = []byte(`{"question":"Is milk an organic compound?"}`)
	req, _ := http.NewRequest("GET", "/upvotequestion", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	a := http.HandlerFunc(database.UpvoteQuestion)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	if resp.Body.String() != `"Upvote Question Successful"` {
		t.Errorf(`Expected message "Upvote Question Successful". Got '%v'`, resp.Body.String())
	}
}
func Test_UpvoteAnswer(t *testing.T) {
	var jsonStr = []byte(`{"Answer":"Uf has good faculty and great research facilities"}`)
	req, _ := http.NewRequest("GET", "/upvoteanswer", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	a := http.HandlerFunc(database.UpvoteAnswer)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	if resp.Body.String() != `"Upvote Answer Successful"` {
		t.Errorf(`Expected message "Upvote Answer Successful". Got '%v'`, resp.Body.String())
	}
}
func Test_DownvoteAnswer(t *testing.T) {
	var jsonStr = []byte(`{"Answer":"Uf has good faculty and great research facilities"}`)
	req, _ := http.NewRequest("GET", "/downvoteanswer", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	a := http.HandlerFunc(database.DownvoteAnswer)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	if resp.Body.String() != `"Downvote Answer Successful"` {
		t.Errorf(`Expected message "Downvote Answer Successful". Got '%v'`, resp.Body.String())
	}
}

func Test_Report(t *testing.T) {
	var jsonStr = []byte(`{"question":"Is milk an organic compound?"}`)
	req, _ := http.NewRequest("GET", "/report", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	a := http.HandlerFunc(database.Report)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	if resp.Body.String() != `"Question Reported Sucessfully"` {
		t.Errorf(`Expected message "Question Reported Sucessfully". Got '%v'`, resp.Body.String())
	}
}
func Test_Report1(t *testing.T) {
	var jsonStr1 = []byte(`{"question":"what is speed of the bullet train?"}`)
	req1, _ := http.NewRequest("GET", "/report", bytes.NewReader(jsonStr1))
	req1.Header.Set("Content-Type", "application/json")
	b := http.HandlerFunc(database.Report)
	resp1 := httptest.NewRecorder()
	b.ServeHTTP(resp1, req1)
	checkResponseCode(t, http.StatusOK, resp1.Code)

	if resp1.Body.String() != `"Question reported Frequently, so it is deleted"` {
		t.Errorf(`Expected message "Question reported Frequently, so it is deleted". Got '%v'`, resp1.Body.String())
	}
}

func Test_GetUnanswered(t *testing.T) {
	req, _ := http.NewRequest("GET", "/getUnanswered", nil)
	req.Header.Set("Content-Type", "application/json")

	a := http.HandlerFunc(database.GetUnanswered)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	if resp.Body.String() != `[[{"Key":"_id","Value":"61fdd6999ff44333f800c14d"},{"Key":"question","Value":"Why me?"},{"Key":"user","Value":"61fdd6999ff44333f800c14c"},{"Key":"upvotes","Value":23},{"Key":"downvotes","Value":1},{"Key":"is_answered","Value":false},{"Key":"followers","Value":[]},{"Key":"topics","Value":[]},{"Key":"answers","Value":[]},{"Key":"comments","Value":[]}],[{"Key":"_id","Value":"61fdf8c7cc1711bfb99ae5f8"},{"Key":"question","Value":"Why Software Engineering?"},{"Key":"user","Value":"61fdf8c7cc1711bfb99ae5f7"},{"Key":"upvotes","Value":100},{"Key":"downvotes","Value":56},{"Key":"is_answered","Value":false},{"Key":"followers","Value":[]},{"Key":"topics","Value":[]},{"Key":"answers","Value":[]},{"Key":"comments","Value":[]}],[{"Key":"_id","Value":"625a30ba2b17f1e955f6672d"},{"Key":"question","Value":"jhhjkjhj"},{"Key":"answer","Value":[]},{"Key":"username","Value":"SU"},{"Key":"downvotes","Value":0},{"Key":"upvotes","Value":0},{"Key":"topic","Value":"jhkjh"},{"Key":"is_answered","Value":false}],[{"Key":"_id","Value":"625a32076a69503f031ce629"},{"Key":"question","Value":"new ques"},{"Key":"answer","Value":[]},{"Key":"username","Value":"SU"},{"Key":"downvotes","Value":0},{"Key":"upvotes","Value":0},{"Key":"topic","Value":"Science"},{"Key":"is_answered","Value":false}],[{"Key":"_id","Value":"625b889e6a69503f031ce62a"},{"Key":"question","Value":"dsdsdsdsdsdsd"},{"Key":"answer","Value":[]},{"Key":"username","Value":"SU"},{"Key":"downvotes","Value":0},{"Key":"upvotes","Value":0},{"Key":"topic","Value":"dsdsdsds"},{"Key":"is_answered","Value":false}],[{"Key":"_id","Value":"62609113d8a699c5c0041145"},{"Key":"question","Value":"what is chipset"},{"Key":"answer","Value":[]},{"Key":"username","Value":"Harshwardhan Chauhan"},{"Key":"downvotes","Value":0},{"Key":"upvotes","Value":0},{"Key":"topic","Value":"Computer"},{"Key":"is_answered","Value":false},{"Key":"report","Value":0}],[{"Key":"_id","Value":"626093e236631a2219f34ca9"},{"Key":"question","Value":"Who was napolean?"},{"Key":"answer","Value":[]},{"Key":"username","Value":"Harshwardhan Chauhan"},{"Key":"downvotes","Value":0},{"Key":"upvotes","Value":0},{"Key":"topic","Value":"History"},{"Key":"is_answered","Value":false},{"Key":"report","Value":0}],[{"Key":"_id","Value":"62609429b9f81a8ed0e6dced"},{"Key":"question","Value":"How are you Praveen?"},{"Key":"answer","Value":[]},{"Key":"username","Value":"Harshwardhan Chauhan"},{"Key":"downvotes","Value":0},{"Key":"upvotes","Value":0},{"Key":"topic","Value":"Health"},{"Key":"is_answered","Value":false},{"Key":"report","Value":0}],[{"Key":"_id","Value":"626094aab9f81a8ed0e6dcee"},{"Key":"question","Value":"Am I just tired or going crazy?"},{"Key":"answer","Value":[]},{"Key":"username","Value":"Harshwardhan Chauhan"},{"Key":"downvotes","Value":0},{"Key":"upvotes","Value":0},{"Key":"topic","Value":"Psychology"},{"Key":"is_answered","Value":false},{"Key":"report","Value":0}],[{"Key":"_id","Value":"6260952ab9f81a8ed0e6dcef"},{"Key":"question","Value":"What is a good investment?"},{"Key":"answer","Value":[]},{"Key":"username","Value":"Harshwardhan Chauhan"},{"Key":"downvotes","Value":0},{"Key":"upvotes","Value":0},{"Key":"topic","Value":"Finance"},{"Key":"is_answered","Value":false},{"Key":"report","Value":0}],[{"Key":"_id","Value":"6260954db9f81a8ed0e6dcf0"},{"Key":"question","Value":"Where is LHC?"},{"Key":"answer","Value":[]},{"Key":"username","Value":"Harshwardhan Chauhan"},{"Key":"downvotes","Value":0},{"Key":"upvotes","Value":0},{"Key":"topic","Value":"Science"},{"Key":"is_answered","Value":false},{"Key":"report","Value":0}],[{"Key":"_id","Value":"626095e5b9f81a8ed0e6dcf1"},{"Key":"question","Value":"How to score good grade in AOA?"},{"Key":"answer","Value":[]},{"Key":"username","Value":"Harshwardhan Chauhan"},{"Key":"downvotes","Value":0},{"Key":"upvotes","Value":0},{"Key":"topic","Value":"Education"},{"Key":"is_answered","Value":false},{"Key":"report","Value":0}],[{"Key":"_id","Value":"62609603b9f81a8ed0e6dcf2"},{"Key":"question","Value":"new question"},{"Key":"answer","Value":[]},{"Key":"username","Value":"Harshwardhan Chauhan"},{"Key":"downvotes","Value":0},{"Key":"upvotes","Value":0},{"Key":"topic","Value":"Cooking"},{"Key":"is_answered","Value":false},{"Key":"report","Value":0}],[{"Key":"_id","Value":"6260b948b9f81a8ed0e6dcf6"},{"Key":"question","Value":"jhjh"},{"Key":"answer","Value":[]},{"Key":"username","Value":"Harshwardhan Chauhan"},{"Key":"downvotes","Value":0},{"Key":"upvotes","Value":0},{"Key":"topic","Value":"Music"},{"Key":"is_answered","Value":false},{"Key":"report","Value":0}]]` {
		t.Errorf(`Expected a list of all unanswered questions and Got '%v'`, resp.Body.String())
	}
}

func Test_SelectedQuestion(t *testing.T) {
	var jsonStr = []byte(`"Question":"why did you choose uf?"`)
	req, _ := http.NewRequest("POST", "/selectedquestion", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	a := http.HandlerFunc(database.SelectedQuestion)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	if resp.Body.String() != `[{"Key":"_id","Value":"6221544f64470d2905d6c503"},{"Key":"answer","Value":"Uf has good ranking and good career fairs"},{"Key":"username","Value":"Nikhil07"},{"Key":"upvotes","Value":81},{"Key":"downvotes","Value":9}][{"Key":"_id","Value":"62216a2594bb3baf476faa07"},{"Key":"answer","Value":"Uf has good faculty and great research facilities"},{"Key":"username","Value":"Loki"},{"Key":"upvotes","Value":76},{"Key":"downvotes","Value":8}][{"Key":"_id","Value":"62216a2594bb3baf476faa08"},{"Key":"answer","Value":"Uf doesn't have location advantage"},{"Key":"username","Value":"Rishab11"},{"Key":"upvotes","Value":20},{"Key":"downvotes","Value":5}]` {
		t.Errorf(`Expected all the data related to the given question. Got '%v'`, resp.Body.String())
	}
}

func Test_AddAnswer(t *testing.T) {
	var jsonStr = []byte(`{"Question": "why did you choose uf?", "Answer": "It was the obvious choice"}`)
	req, _ := http.NewRequest("POST", "/addAnswer", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	user := database.User{"Harshwardha Chauhan", "harshwardhan0812@gmail.com", "SU", "password", []string{}, 0, 0, []primitive.ObjectID{}, []primitive.ObjectID{}}

	ctxWithUser := context.WithValue(req.Context(), 0, &user)
	reqWithUser := req.WithContext(ctxWithUser)
	a := http.HandlerFunc(database.AddAnswer)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, reqWithUser)
	checkResponseCode(t, http.StatusOK, resp.Code)
	if resp.Body.String() != `"Answer succesfully added."` {
		t.Errorf(`Expected "Answer succesfully added." as response. Got '%v'`, resp.Body.String())
	}
}

func Test_SetUserCategory(t *testing.T) {
	var jsonStr = []byte(`{"Topic": ["Science", "Sports"]}`)
	req, _ := http.NewRequest("POST", "/setUserCategory", bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	user := database.User{"Harshwardha Chauhan", "harshwardhan0812@gmail.com", "SU", "password", []string{}, 0, 0, []primitive.ObjectID{}, []primitive.ObjectID{}}

	ctxWithUser := context.WithValue(req.Context(), 0, &user)
	reqWithUser := req.WithContext(ctxWithUser)
	a := http.HandlerFunc(database.SetUserCategory)
	resp := httptest.NewRecorder()
	a.ServeHTTP(resp, reqWithUser)
	checkResponseCode(t, http.StatusOK, resp.Code)
	if resp.Body.String() != `"Update Successful"` {
		t.Errorf(`Expected "Update Successful" as response. Got '%v'`, resp.Body.String())
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
