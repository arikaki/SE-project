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
	"golang.org/x/crypto/bcrypt"
)

func BsonUser(fullname string, email string, username string, password string, followers []primitive.ObjectID,
	following []primitive.ObjectID, topics []string, question []primitive.ObjectID, answer []primitive.ObjectID) bson.D {
	return bson.D{
		{"fullname", fullname},
		{"email", email},
		{"username", username},
		{"password", password},
		{"followers", followers},
		{"following", following},
		{"topics", topics},
		{"question", question},
		{"answer", answer},
	}
}
func BsonQuestion(question string /* upvotes int, comments []primitive.ObjectID*/, answer []primitive.ObjectID, username string, downvotes int) bson.D {
	return bson.D{
		{"question", question},
		{"answer", answer},
		{"username", username},
		{"downvotes", downvotes},
	}

}
func BsonAnswer(answer string, username string, upvotes int, downvotes int) bson.D {
	return bson.D{
		{"answer", answer},
		{"username", username},
		{"upvotes", upvotes},
		{"downvotes", downvotes},
	}

}

type User struct {
	Fullname  string               `json:"Fullname"`
	Email     string               `json:"Email"`
	Username  string               `json:"UserName"`
	Password  string               `json:"Password"`
	Topics    []string             `json:"Topics"`
	Upvotes   int                  `json:"Upvotes"`
	Downvotes int                  `json:"Downvotes"`
	Question  []primitive.ObjectID `json:"Questions"`
	Answer    []primitive.ObjectID `json:"Answer"`
}

type Question []primitive.ObjectID

func getUserCollection() *mongo.Collection {
	db, dbPresent := os.LookupEnv("DBName")
	if !dbPresent {
		db = "KoraDB"
	}
	var collection = client.Database(db).Collection("Users")
	return collection
}

func FetchUser(username string) User {
	collection := getUserCollection()
	var getResult bson.D
	err := collection.FindOne(context.TODO(), bson.D{
		{"username", bson.D{{"$eq", username}}},
	}).Decode((&getResult))
	if err != nil {
		fmt.Println("ERROR", err)
	}
	var fetchedUser User
	json.Marshal(getResult)
	fmt.Println("ID", getResult)
	bsonBytes, _ := bson.Marshal(getResult)
	bson.Unmarshal(bsonBytes, &fetchedUser)
	return fetchedUser
}

func InsertUsers(w http.ResponseWriter, r *http.Request) {
	var post User

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println("error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(post.Password), 14)
	user := BsonUser(post.Fullname, post.Email, post.Username, string(password), []primitive.ObjectID{}, []primitive.ObjectID{}, post.Topics, []primitive.ObjectID{}, []primitive.ObjectID{})

	collection := getUserCollection()
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertResult.InsertedID)
}
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(0).(*User)
	jsonUser, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUser)
}
func GetUser(userName string) (*User, error) {
	collection := getUserCollection()
	var getResult bson.D

	err := collection.FindOne(context.TODO(), bson.D{
		{"username", bson.D{{"$eq", userName}}},
	}).Decode((&getResult))
	fmt.Println(getResult)
	if err != nil {
		fmt.Println("ERROR", err)
		return nil, err
	}
	var fetchedUser User

	bsonBytes, _ := bson.Marshal(getResult)
	bson.Unmarshal(bsonBytes, &fetchedUser)
	return &fetchedUser, nil
}

func InsertDummyUser(w http.ResponseWriter, r *http.Request) {
	collection := getUserCollection()
	insertManyResult, err := collection.InsertMany(context.TODO(), DummyData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

func InsertDummyQuestion(w http.ResponseWriter, r *http.Request) {
	collection := getUserCollection()
	insertManyResult, err := collection.InsertMany(context.TODO(), DummyQuestion)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

}
func InsertDummyAnswer(w http.ResponseWriter, r *http.Request) {
	collection := getUserCollection()
	insertManyResult, err := collection.InsertMany(context.TODO(), DummyAnswer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

}

// bson.D{
// 	{"fullname", "Harshwardhan"},
// 	{"email", "harshwardhan0812@gmail.com"},
// 	{"username", "SU"},
// 	{"password", "password"},
// {"followers", [...]primitive.ObjectID{primitive.ObjectID({"$oid": "61fc48c0188132eecabf661e"}), primitive.ObjectID({"$oid": "61fc48c0188132eecabf661f"}),
// 	primitive.ObjectID({"$oid": "61fc48c0188132eecabf6620"}), primitive.ObjectID({"$oid": "61fc48c0188132eecabf6621"})}},
// {"following", [...]primitive.ObjectID{primitive.ObjectID("61fc48c0188132eecabf661e"), primitive.ObjectID("61fc48c0188132eecabf661f"),
// primitive.ObjectID("61fc48c0188132eecabf6620"), primitive.ObjectID("61fc48c0188132eecabf6621")}},
// {"topics", [...]primitive.ObjectID{primitive.ObjectID("61fc48c0188132eecabf661e"), primitive.ObjectID("61fc48c0188132eecabf661f"),
// primitive.ObjectID("61fc48c0188132eecabf6620"), primitive.ObjectID("61fc48c0188132eecabf6621")}},
// }
