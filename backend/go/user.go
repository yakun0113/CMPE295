package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func signUp(c *gin.Context) {
	var user User
	var dbUser User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	collection := client.Database("octopusDB").Collection("user")
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)
	if err == nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "This email is already registered!"})
		return
	}
	user.Password = getHash([]byte(user.Password))
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else if result != nil {
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "Signed up successfully!"})
	}

}

func signIn(c *gin.Context) {

	var user User
	var dbUser User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	collection := client.Database("octopusDB").Collection("user")
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)

	if err != nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Account not registered!"})
		return
	}
	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Wrong password!"})
		return
	}
	jwtToken, err := GenerateJWT(dbUser.Email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Something's wrong, please try again!"})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"token": jwtToken, "message": "Welcome " + dbUser.Name + "!"})

}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
