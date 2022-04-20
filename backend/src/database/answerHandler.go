package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func BsonAnswer(answer string, username string, upvotes int, downvotes int) bson.D {
	return bson.D{
		{"answer", answer},
		{"username", username},
		{"upvotes", upvotes},
		{"downvotes", downvotes},
	}
}

// type searchPost struct {
// 	Search string `json:"Search"`
// }

type answerData struct {
	Answer    string `json:"Answer"`
	Question  string `json:"Question"`
	Upvotes   int    `json:"Upvotes"`
	Downvotes int    `json:"Downvotes"`
}

func getAnswerCollection() *mongo.Collection {
	db, dbPresent := os.LookupEnv("DBName")
	if !dbPresent {
		db = "KoraDB"
	}
	var collection = client.Database(db).Collection("Answers")
	return collection
}

func AddAnswer(w http.ResponseWriter, r *http.Request) {
	var post answerData

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println("error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := r.Context().Value(0).(*User)

	ansColl := getAnswerCollection()
	quesColl := getQuestionCollection()

	answerBson := BsonAnswer(post.Answer, user.Username, 0, 0)
	insertResult, err := ansColl.InsertOne(context.TODO(), answerBson)
	if err != nil {
		log.Fatal(err)
	}
	result, err := quesColl.UpdateOne(
		context.TODO(),
		bson.M{"question": post.Question},
		bson.D{
			{"$set", bson.D{{"is_answered", true}}},
			{"$push", bson.D{{"answer", insertResult.InsertedID}}},
		},
	)
	if err != nil {
		log.Fatal(err, result)
	}

	jsonResponse, err := json.Marshal("Answer succesfully added.")
	if err != nil {
		return
	}
	w.Write(jsonResponse)

}
