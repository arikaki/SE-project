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

func BsonQuestion(question string, answer []primitive.ObjectID, username string, downvotes int, upvotes int, topic string, is_answered bool, report int) bson.D {
	return bson.D{
		{"question", question},
		{"answer", answer},
		{"username", username},
		{"downvotes", downvotes},
		{"upvotes", upvotes},
		{"topic", topic},
		{"is_answered", is_answered},
		{"report", report},
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
	Topic      string `json:"Topic"`
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

	questionBson := BsonQuestion(post.Question, []primitive.ObjectID{}, user.Username, post.Downvotes, post.Upvotes, post.Topic, false, 0)
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
func UpvoteQuestion(w http.ResponseWriter, r *http.Request) {
	collection := getQuestionCollection()
	var data selectedQuestion
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	filter := bson.D{{"question", data.Question}}
	update := bson.D{{"$inc", bson.D{{"upvotes", 1}}}}
	result, err1 := collection.UpdateOne(context.Background(), filter, update)
	if err1 != nil {
		//
	}
	fmt.Println(result.ModifiedCount)
	jsonResponse, err := json.Marshal("Upvote Question Successful")
	if err != nil {
		return
	}
	w.Write(jsonResponse)

}

func DownvoteQuestion(w http.ResponseWriter, r *http.Request) {
	collection := getQuestionCollection()
	var data selectedQuestion
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	filter := bson.D{{"question", data.Question}}
	update := bson.D{{"$inc", bson.D{{"downvotes", -1}}}}
	result, err1 := collection.UpdateOne(context.Background(), filter, update)
	if err1 != nil {
		log.Fatal(err)
	}
	fmt.Println(result.ModifiedCount)
	jsonResponse, err := json.Marshal("Downvote Question Successful")
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func Report(w http.ResponseWriter, r *http.Request) {
	collection := getQuestionCollection()
	var data selectedQuestion
	// var rep reprt
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	filter := bson.D{{"question", data.Question}}
	projection := bson.D{{"report", 1}}
	opts := options.FindOne().SetProjection(projection)
	var result bson.D
	err1 := collection.FindOne(context.TODO(), filter, opts).Decode(&result)
	if err1 != nil {
		//
	}
	// bsonBytes, _ := bson.Marshal(result)
	// bson.Unmarshal(bsonBytes, &rep)
	var match int32 = 4
	if result[1].Value != match {
		filter := bson.D{{"question", data.Question}}
		update := bson.D{{"$inc", bson.D{{"report", 1}}}}
		result, err1 := collection.UpdateOne(context.Background(), filter, update)
		if err1 != nil {
			//
		}
		fmt.Println(result.ModifiedCount)
		jsonResponse, err := json.Marshal("Question Reported Sucessfully")
		if err != nil {
			return
		}
		w.Write(jsonResponse)
	} else {
		result, err := collection.DeleteOne(r.Context(), bson.D{{"question", data.Question}})
		if err1 != nil {
			log.Fatal(err)
		}
		if result.DeletedCount == 0 {
			fmt.Println("Question not found.")
		}
		jsonResponse, err := json.Marshal("Question reported Frequently, so it is deleted")
		if err != nil {
			return
		}
		w.Write(jsonResponse)
	}
}
func TopQuestion(w http.ResponseWriter, r *http.Request) {
	collection := getQuestionCollection()
	var data topic
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// coll := getUserCollection()
	// var user *User
	// user = r.Context().Value(0).(*User)
	// project := bson.D{{"topic", 0}}
	// opts := options.FindOne().SetProjection(project)

	filter := bson.D{{"topic", bson.D{{"$in", data.Topic}}}}
	// // sort := bson.D{{"upvotes", -1}}
	projection := bson.D{{"question", 1}, {"_id", 0}, {"topic", 1}}
	opts := options.Find().SetProjection(projection)
	var result bson.D
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {

	}
	for cursor.Next(context.TODO()) {
		if err1 := cursor.Decode(&result); err1 != nil {
			log.Fatal(err1)
		}
		jsonUser, _ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonUser)
	}

}

func SelectedQuestion(w http.ResponseWriter, r *http.Request) {
	collection := getAnswerCollection()
	coll := getQuestionCollection()
	var data selectedQuestion
	var ans answer
	// var returnAnswer returnAns
	// var results []returnAns
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// // project := bson.D{{"password", 0}}
	// // opts := options.FindOne().SetProjection(project)

	// // err := collection.FindOne(context.TODO(), bson.D{
	// // 	{"username", bson.D{{"$eq", userName}}},
	// // }, opts).Decode((&getResult))

	filter := bson.D{{"question", data.Question}}
	projection := bson.D{{"answer", 1}}
	opts := options.FindOne().SetProjection(projection)
	var result bson.D
	err1 := coll.FindOne(context.TODO(), filter, opts).Decode(&result)
	if err1 != nil {
		//
	}
	// result1, _ := json.Marshal(result)
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &ans)

	filter1 := bson.D{{"_id", bson.D{{"$in", ans.Answer}}}}
	projection1 := bson.D{{"username", 1}, {"answer", 1}, {"upvotes", 1}, {"downvotes", 1}}
	opt := options.Find().SetProjection(projection1)
	cursor, err := collection.Find(context.TODO(), filter1, opt)
	if err != nil {
		//
	}
	var result1 bson.D
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&result1)
		if err != nil {
			log.Fatal(err)
		} //  else {
		// 	// answerList, _ := json.Marshal(result1)
		// 	// fmt.Println("results", result1)
		// 	// bson.Unmarshal(answerList, &returnAnswer)
		// 	// results = append(results, result1)
		// 	fmt.Println("results")
		// }
		jsonUser, _ := json.Marshal(result1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonUser)
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// w.Write(results)
	}
}
