package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// type dataFormat struct {
// 	name string
// 	email string
// 	username string
// 	passworrd string

// }

var p, _ = bcrypt.GenerateFromPassword([]byte("password"), 14)
var password = string(p)

var DummyData = []interface{}{
	BsonUser("Harshwardhan", "harshwardhan0812@gmail.com", "SU", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Lokesh", "Lokesh248@gmail.com", "Loki", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Nikhil", "Nikhil99@gmail.com", "Nikhil07", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Rishab", "Rishab21@gmail.com", "Rishab11", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Anudeep", "Anudeep8@gmail.com", "Bar", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Srisai", "Srisai8@gmail.com", "Srisai9", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Bhargav", "Bhargav12@gmail.com", "Bhargav12", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Sagar", "Sagar32@gmail.com", "Sagar9", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Kiran", "Kiran212@gmail.com", "Kirank5", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Chaitanya", "Chaitanya28@gmail.com", "Chaitanya11", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Aakash", "Aakash24@gmail.com", "Aakash78", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Vishal", "Vishal14@gmail.com", "Vishal15", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Jagan", "jagan02@gmail.com", "jagan8", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Nitish", "Aakash24@gmail.com", "Nitish12", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Praveen", "praveen18@gmail.com", "Praveen18", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Vamsi", "Vamsi62@gmail.com", "Vamsi09", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("yashas", "yashas2@gmail.com", "yashas17", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Sai Ram reddy", "sairam79@gmail.com", "Sairam", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Sai reddy", "saireddy9@gmail.com", "Sai11", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Nitin", "Nitin11@gmail.com", "Nitin36", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Anil", "Anil27@gmail.com", "Anil23", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Amit", "amit18@gmail.com", "Amit3", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("kalyan", "kalyan11@gmail.com", "kalyan09", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("srikar", "srikar56@gmail.com", "srikar05", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Riyaz", "Riyaz02@gmail.com", "Riyaz44", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Harsha", "Harsha14@gmail.com", "Harsha4", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Dheeraj", "Dheeraj@gmail.com", "Dheeraj89", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Vikram", "Vikram52@gmail.com", "Vikram84", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Nandigam", "Nandigam54@gmail.com", "Nandigam04", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("arif", "arif11@gmail.com", "arif4", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Nagarjuna", "Nagarjuna55@gmail.com", "Nag11", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Pawan Kalyan", "pkalyan@gmail.com", "PK", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Chiranjeevi", "Chiranjeevi55@gmail.com", "Chiru", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Prabhas", "Prabhas27@gmail.com", "Prabhas5", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Ajith", "Ajith77@gmail.com", "ajith", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Rajinikanth", "Rajinikant36@gmail.com", "Thalaiva", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Sharwanand", "Sharwa11@gmail.com", "sharwa", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("NTRamarao", "NTR11@gmail.com", "NTR", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Ravi teja", "Raviteja@gmail.com", "RTR", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("salman bhai", "salman99@gmail.com", "Sallubhai", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Shahrukh", "srk11@gmail.com", "SRK12", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Katrina", "katrina24@gmail.com", "katkaif", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Kajal", "kajal27@gmail.com", "Kajal", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Samantha", "samantha@gmail.com", "Sam", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Alluarjun", "Alluarjun99@gmail.com", "Bunny", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Robertdowney", "Robertjr99@gmail.com", "Ironman", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Chris Hemsworth", "ChrisHems@gmail.com", "Thor", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Scarlett", "Scarlett9@gmail.com", "Blackwidow", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Messi", "Lmessi@gmail.com", "Messi10", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Chrisevans", "Chrisevans@gmail.com", "Captainamerica", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Sachin Tendulkar", "SRT@gmail.com", "Sachin", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Rohit", "Rohit45@gmail.com", "Hitman", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Virat", "VK07@gmail.com", "Vkohli", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Devilliers", "ABD360@gmail.com", "ABD", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
	BsonUser("Teja", "Teja18@gmail.com", "Teja12", "password", []primitive.ObjectID{}, []primitive.ObjectID{}, []string{"Technology", "Art", "Sports"}, []primitive.ObjectID{}, []primitive.ObjectID{}),
}
var DummyQuestion = []interface{}{
	BsonQuestion("why did you choose uf?", []primitive.ObjectID{}, "Chiru", 0, 9, "Education", 0),
	BsonQuestion("How can modern technology help evolve?", []primitive.ObjectID{}, "SU", 2, 5, "Technology", 1),
	BsonQuestion("What’s a good investment for 2022?", []primitive.ObjectID{}, "Loki", 9, 18, "Investment", 1),
	BsonQuestion("What does “G.O.A.T” mean?", []primitive.ObjectID{}, "Nikhil07", 1, 10, "Sports", 1),
	BsonQuestion("Is milk an organic compound?", []primitive.ObjectID{}, "srikar05", 3, 6, "Chemistry", 2),
	BsonQuestion("How long are GRE scores valid?", []primitive.ObjectID{}, "Bar", 11, 8, "Education", 1),
	BsonQuestion("What does “absolute refractive index of glass is 1.5” mean?", []primitive.ObjectID{}, "Srisai9", 8, 3, "Technology", 0),
	BsonQuestion("Where is Gainesville located?", []primitive.ObjectID{}, "Kalyan09", 11, 11, "General Knowledge", 0),
	BsonQuestion("Which is the worlds number one cricket team in T20's?", []primitive.ObjectID{}, "Nandigam04", 0, 8, "Sports", 1),
	BsonQuestion("What is the worlds most spoken language?", []primitive.ObjectID{}, "Bunny", 5, 5, "General Knowledge", 0),
	BsonQuestion("What does “something being out of” mean?", []primitive.ObjectID{}, "Nag11", 20, 6, "General Knowledge", 0),
	BsonQuestion("How many egg whites should one eat per day?", []primitive.ObjectID{}, "Amit3", 4, 10, "Health", 0),
	BsonQuestion("What are examples of demand-side market failures?", []primitive.ObjectID{}, "Nitin36", 6, 6, "Investment", 1),
	BsonQuestion("What is super bowl?", []primitive.ObjectID{}, "Sachin", 10, 15, "Sports", 0),
	BsonQuestion("Who is the owner of tesla?", []primitive.ObjectID{}, "Vamsi09", 4, 10, "General Knowledge", 0),
	BsonQuestion("what is the speed of the bullet train?", []primitive.ObjectID{}, "Nitish12", 18, 8, "Technology", 0),
}
var DummyAnswer = []interface{}{
	BsonAnswer("Uf has good ranking and career fairs", "Nikhil07", 81, 9),
	BsonAnswer("Uf has good faculty and great research facilities", "Loki", 72, 12),
	BsonAnswer("Uf doesn't have location advantage", "Rishab11", 20, 5),
	BsonAnswer("The best investment you can make is in share market, properties, gold", "Bhargav12", 54, 4),
	BsonAnswer("“G.O.A.T” means Greatest of all time", "Bar", 79, 6),
	BsonAnswer("Technology has helped us to fly, drive, sail,communicate", "Srisai9", 20, 5),
	BsonAnswer("It’s a mixture of mainly organic compounds.", "Anil23", 57, 15),
	BsonAnswer("They are valid for a period of five years", "yashas17", 28, 9),
	BsonAnswer(" if the refractive index of glass is 1.5, it means the medium glass is 1.5 times denser than air or conversely the speed of light in glass is 1.5 times greater than in air", "Praveen18", 68, 20),
	BsonAnswer("It is located in Florida, US", "Sagar9", 37, 8),
	BsonAnswer("Indian Team stands in first place according to current rankings", "Chaitanya11", 45, 5),
	BsonAnswer("English is the most spoken language in th world", "Aakash78", 60, 12),
	BsonAnswer("When something ran out, it means it is finished,To finish or use something until there is none left.", "Nitish12", 33, 9),
	BsonAnswer("There is no recommended limit to eat white eggs.", "Harsha4", 46, 20),
	BsonAnswer("One would be the recent fashion failure of real fur for womens’ outerwear, as that cohort became animal-cruelty-conscious and began to opt for faux fur.", "jagan8", 32, 5),
	BsonAnswer("The Super Bowl is the annual playoff championship game of the National Football League. It has served as the final game of every NFL season", "Dheeraj89", 19, 2),
	BsonAnswer("Elon Musk co-founded and is the CEO of Tesla", "Kalyan09", 56, 15),
	BsonAnswer("The bullet train is capable of reaching a maximum speed of 320kms per hour, the bullet train offers riders an exceptionally unique and efficient travel experience.", "srikar05", 64, 6),
}
