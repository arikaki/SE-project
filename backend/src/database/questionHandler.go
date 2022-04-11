package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func bsonQuestion(question string, user primitive.ObjectID, upvotes int, downvotes int, is_answered bool, followers []primitive.ObjectID,
	topics []primitive.ObjectID, answers []primitive.ObjectID, comments []primitive.ObjectID) bson.D {
	return bson.D{
		{"question", question},
		{"user", user},
		{"upvotes", upvotes},
		{"downvotes", downvotes},
		{"is_answered", is_answered},
		{"followers", followers},
		{"topics", topics},
		{"answers", answers},
		{"comments", comments},
	}
}

type searchPost struct {
	Search string `json:"Search"`
}

type Post struct {
	Question   string `json:"Question"`
	Isanswered bool   `json:"Isanswered"`
	Upvotes    int    `json:"Upvotes"`
	Downvotes  int    `json:"Downvotes"`
}

func getQuestionCollection() *mongo.Collection {
	db, dbPresent := os.LookupEnv("DBName")
	if !dbPresent {
		db = "KoraDB"
	}
	var collection = client.Database(db).Collection("Questions")
	return collection
}

func AskQ(w http.ResponseWriter, r *http.Request) {
	var post Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println("error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := r.Context().Value(0).(*User)
	fmt.Println("logged user from middleware", user)

	collection := getQuestionCollection()

	questionBson := bsonQuestion(post.Question, primitive.NewObjectID(), post.Upvotes, post.Downvotes, false, []primitive.ObjectID{}, []primitive.ObjectID{},
		[]primitive.ObjectID{}, []primitive.ObjectID{})
	// insertResult, err := collection.InsertOne(context.TODO(), harshwardhan)
	insertResult, err := collection.InsertOne(context.TODO(), questionBson)
	if err != nil {
		log.Fatal(err)
	}
	jsonResponse, err := json.Marshal("Question succesfully added.")
	if err != nil {
		return
	}
	fmt.Println("Inserted document: ", insertResult.InsertedID)
	w.Write(jsonResponse)

}

func GetAllQ(w http.ResponseWriter, r *http.Request) {
	var allQuestions []*Post
	collection := getQuestionCollection()
	result, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for result.Next(context.TODO()) {
		var fetchedQuestion Post
		err := result.Decode(&fetchedQuestion)
		if err != nil {
			// return allQuestions, err
		}

		allQuestions = append(allQuestions, &fetchedQuestion)
	}

	if err := result.Err(); err != nil {
		// return allQuestions, err
	}

	// once exhausted, close the cursor
	result.Close(context.TODO())

	if len(allQuestions) == 0 {
		fmt.Println("return", allQuestions)
		// return allQuestions, mongo.ErrNoDocuments
	}

	// return tasks, nil

	fmt.Println("Fetched Question", allQuestions)
}

func FindMatchingQuestions(w http.ResponseWriter, r *http.Request) {
	var data searchPost

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	collection := getQuestionCollection()
	model := mongo.IndexModel{Keys: bson.D{{"question", "text"}}}
	name, err := collection.Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created: " + name)
	filter := bson.D{{"$text", bson.D{{"$search", data.Search}}}}
	projection := bson.D{{"question", 1}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	var results []bson.D
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(context.TODO())
	} else {
		for cursor.Next(context.TODO()) {
			var result bson.D
			err := cursor.Decode(&result)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			} else {
				results = append(results, result)
			}
		}
	}
	jsonResponse, err := json.Marshal(results)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func GetUnanswered(w http.ResponseWriter, r *http.Request) {
	collection := getQuestionCollection()
	filter := bson.D{{"is_answered", bson.D{{"$eq", false}}}}
	cursor, err := collection.Find(context.TODO(), filter)
	var results []bson.D
	if err != nil {
		fmt.Println("Finding unanswered questions ERROR:", err)
		defer cursor.Close(context.TODO())
	} else {
		for cursor.Next(context.TODO()) {
			var result bson.D
			err := cursor.Decode(&result)
			fmt.Println("cursor and result", cursor, result)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			} else {
				results = append(results, result)
			}
		}
	}
	jsonResponse, err := json.Marshal(results)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}
