package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client
var SECRET_KEY = []byte("cmpe295")

func connectDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Brian80433:cmpe295@cluster0.oolyk.mongodb.net/octopus?retryWrites=true&w=majority"))
}

type User struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func signUp(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var user User
	var dbUser User

	collection := client.Database("octopus").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	json.NewDecoder(request.Body).Decode(&user)
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)

	if err == nil {
		response.Write([]byte("This email is already registered!"))
		return
	}

	user.Password = getHash([]byte(user.Password))
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
	} else if result != nil {
		response.Write([]byte("Signed up successfully!"))
	}
}

func signIn(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var user User
	var dbUser User

	collection := client.Database("octopus").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(request.Body).Decode(&user)
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)

	if err != nil {
		response.Write([]byte("Account not registered!"))
		return
	}
	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		response.Write([]byte("Wrong password!"))
		return
	}
	userName := dbUser.Name
	jwtToken, err := GenerateJWT()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	response.Write([]byte(`{"token":"` + jwtToken + `", "name":"` + userName + `"}`))

}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
