package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func BsonUser(fullname string, email string, username string, password string, followers []primitive.ObjectID,
	following []primitive.ObjectID, topics []primitive.ObjectID) bson.D {
	return bson.D{
		{"fullname", fullname},
		{"email", email},
		{"username", username},
		{"password", password},
		{"followers", followers},
		{"following", following},
		{"topics", topics},
	}
}

type User struct {
	Fullname string `json:"Fullname"`
	Email    string `json:"Email"`
	Username string `json:"UserName"`
	Password string `json:"Password"`
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
	user := BsonUser(post.Fullname, post.Email, post.Username, string(password), []primitive.ObjectID{}, []primitive.ObjectID{}, []primitive.ObjectID{})

	var collection = client.Database("KoraDB").Collection("Users")
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertResult.InsertedID)
}

func GetUser(userName string) (*User, error) {
	var collection = client.Database("KoraDB").Collection("Users")
	var getResult bson.D
	err := collection.FindOne(context.TODO(), bson.D{
		{"username", bson.D{{"$eq", userName}}},
	}).Decode((&getResult))
	if err != nil {
		fmt.Println("ERROR", err)
		return nil, err
	}
	var fetchedUser User

	bsonBytes, _ := bson.Marshal(getResult)
	bson.Unmarshal(bsonBytes, &fetchedUser)
	return &fetchedUser, nil
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
