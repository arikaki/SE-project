package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type login struct {
	Email    string `json:"Email"`
	Username string `json:"UserName"`
	Password string `json:"Password"`
}

var SecretKey = []byte(os.Getenv("Cookie_Key"))

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

type selectedQuestion struct {
	Question string `json:"Question"`
}
type selectedAnswer struct {
	Answer string `json:"Answer"`
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

type topic struct {
	Topic []string `json:"Topic"`
}

type answer struct {
	Answer []primitive.ObjectID `json:Answer"`
}

// type reprt struct {
// 	Report int `json:Report"`
// }
type userName struct {
	Username string `json:"UserName"`
}

type returnAns struct {
	Username  string `json:"Username"`
	Answers   string `json:"Answer"`
	Upvotes   int    `json:"Upvotes"`
	Downvotes int    `json:"Downvotes"`
}

func getUserCollection() *mongo.Collection {
	db, dbPresent := os.LookupEnv("DBName")
	if !dbPresent {
		db = "KoraDB"
	}
	var collection = client.Database(db).Collection("Users")
	return collection
}

// func getQuesCollection() *mongo.Collection {
// 	db, dbPresent := os.LookupEnv("DBName")
// 	if !dbPresent {
// 		db = "KoraDB"
// 	}
// 	var QuestionCollection = client.Database(db).Collection("Questions")
// 	return QuestionCollection
// }
// func getAnsCollection() *mongo.Collection {
// 	db, dbPresent := os.LookupEnv("DBName")
// 	if !dbPresent {
// 		db = "KoraDB"
// 	}
// 	var AnsCollection = client.Database(db).Collection("Answers")
// 	return AnsCollection
// }

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

func loginUser(w http.ResponseWriter, username string, password string, isRegister bool) {
	fetchedUser, _ := GetUser(username)
	if err := bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(password)); err != nil {
		//err
	}
	claims := &Claims{
		Username: fetchedUser.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		// err
	}

	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{Name: "Session", Value: tokenString, Expires: expiration}
	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	//convert struct to JSON
	response := "Login Successful"
	if isRegister {
		response = "Inserted User"
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	//update response
	w.Write(jsonResponse)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	loginUser(w, data.Username, data.Password, false)
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
	_ = insertResult
	if err != nil {
		log.Fatal(err)
	}
	loginUser(w, post.Username, post.Password, true)
}

func GetUser(userName string) (*User, error) {
	collection := getUserCollection()
	var getResult bson.D
	project := bson.D{{"password", 0}}
	opts := options.FindOne().SetProjection(project)

	err := collection.FindOne(context.TODO(), bson.D{
		{"username", bson.D{{"$eq", userName}}},
	}, opts).Decode((&getResult))
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
	var user *User
	user = r.Context().Value(0).(*User)
	// usersC := Collection("users")
	//  ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	// var getResult bson.D
	// project := bson.D{{"password", 0}}
	// opts := options.FindOne().SetProjection(project)

	// err := collection.FindOne(context.TODO(), bson.D{
	// 	{"username", bson.D{{"$eq", userName}}},
	// }, opts).Decode((&getResult))
	result, err := collection.DeleteOne(r.Context(), bson.D{{"username", bson.D{{"$eq", user.Username}}}})
	if err != nil {
		fmt.Println("failed to delete the user", err)
	}
	if result.DeletedCount == 0 {
		fmt.Println("user not found.")
	}
	jsonResponse, err := json.Marshal("User is deleted")
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

// jsonResponse, err := json.Marshal(results)
// if err != nil {
// 	return
// }
// w.Write(jsonResponse)

// fmt.Println("answers", rep)
// fmt.Println("Data", data)
// fmt.Println("result", result)
// jsonUser, _ := json.Marshal(result)
// w.Header().Set("Content-Type", "application/json")
// w.WriteHeader(http.StatusOK)
// w.Write(jsonUser)

// }

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
