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
func BsonQuestion(question string /* upvotes int, comments []primitive.ObjectID*/, answer []primitive.ObjectID, username string, downvotes int, upvotes int, topic string) bson.D {
	return bson.D{
		{"question", question},
		{"answer", answer},
		{"username", username},
		{"downvotes", downvotes},
		{"upvotes", upvotes},
		{"topic", topic},
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

type topQuestion struct {
	Topics []string `json:"Topics"`
}
type selectedQuestionId struct {
	selectedQuestionId primitive.ObjectID `json:"Selectedquestion"`
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

type userName struct {
	Username string `json:"UserName"`
}

func getUserCollection() *mongo.Collection {
	db, dbPresent := os.LookupEnv("DBName")
	if !dbPresent {
		db = "KoraDB"
	}
	var collection = client.Database(db).Collection("Users")
	return collection
}
func getQuesCollection() *mongo.Collection {
	db, dbPresent := os.LookupEnv("DBName")
	if !dbPresent {
		db = "KoraDB"
	}
	var QuestionCollection = client.Database(db).Collection("Questions")
	return QuestionCollection
}

func getAnsCollection() *mongo.Collection {
	db, dbPresent := os.LookupEnv("DBName")
	if !dbPresent {
		db = "KoraDB"
	}
	var AnswerCollection = client.Database(db).Collection("Answers")
	return AnswerCollection
}

func FetchUser(w http.ResponseWriter, r *http.Request) {
	var data userName
	var user *User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if data.Username == "" {
		user = r.Context().Value(0).(*User)
	} else {
		user, err = GetUser(data.Username)
		if err != nil {
			http.Error(w, "User Not Found", http.StatusForbidden)
			return
		}
	}
	jsonUser, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUser)
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

func GetUser(userName string) (*User, error) {
	collection := getUserCollection()
	var getResult bson.D
	project := bson.D{{"password", 0}}
	opts := options.FindOne().SetProjection(project)

	err := collection.FindOne(context.TODO(), bson.D{
		{"username", bson.D{{"$eq", userName}}},
	}, opts).Decode((&getResult))
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

//Delete user from DB
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	collection := getUserCollection()
	user := r.Context().Value(0).(*User)
	result, err := collection.DeleteOne(r.Context(), bson.D{
		{"username", bson.D{{"$eq", user.Username}}},
	})
	fmt.Println("User is deleted", result)
	if err != nil {
		fmt.Println("failed to delete the user", err)
	}
	if result.DeletedCount == 0 {
		fmt.Println("user not found.")
	}
}
func TopQuestion(w http.ResponseWriter, r *http.Request) {
	collection := getQuesCollection()
	var data topQuestion

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.D{{"topic", bson.D{{"$in", data.Topics}}}}
	// sort := bson.D{{"upvotes", -1}}
	projection := bson.D{{"question", 1}, {"_id", 0}, {"topic", 1}}
	opts := options.Find().SetProjection(projection)
	var result bson.D
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		//
	}
	for cursor.Next(context.TODO()) {
		if err1 := cursor.Decode(&result); err1 != nil {
			log.Fatal(err1)
		}
		fmt.Println(result)
		jsonUser, _ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonUser)
	}

}

func SelectedQuestion(w http.ResponseWriter, r *http.Request) {
	collection := getAnsCollection()
	coll := getQuesCollection()
	var data selectedQuestionId

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// project := bson.D{{"password", 0}}
	// opts := options.FindOne().SetProjection(project)

	// err := collection.FindOne(context.TODO(), bson.D{
	// 	{"username", bson.D{{"$eq", userName}}},
	// }, opts).Decode((&getResult))

	filter := bson.D{{"_id", data.selectedQuestionId}}
	projection := bson.D{{"answer", 1}}
	opts := options.FindOne().SetProjection(projection)
	var result bson.D
	err1 := coll.FindOne(context.TODO(), filter, opts).Decode(&result)

	filter := bson.D{{"answer", bson.D{{"$in", answer}}}}
	projection := bson.D{{"answer", 1}, {"username", 1}, {"upvotes", 1}, {"downvotes", 1}}
	opt := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opt)
	if err != nil {
		//
	}
	for cursor.Next(context.TODO()) {
		if err1 := cursor.Decode(&result); err1 != nil {
			log.Fatal(err1)
		}
		fmt.Println(result)
		jsonUser, _ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonUser)
	}

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
