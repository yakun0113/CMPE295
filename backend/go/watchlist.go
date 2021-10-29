package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func addItem(c *gin.Context) {

	var item Product

	if err := c.BindJSON(&item); err != nil {
		return
	}
	email := c.Param("email")
	collection := client.Database("octopusDB").Collection("user")
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"email": email},
		bson.M{"$push": bson.M{"watchlist": item}})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please try again!"})
	} else if result != nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Added to watchlist!"})

	}

}

func getItems(c *gin.Context) {
	var dbUser User
	email := c.Param("email")
	collection := client.Database("octopusDB").Collection("user")
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&dbUser)

	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{"result": dbUser.Watchlist})
	}
}

func deleteItem(c *gin.Context) {
	email := c.Param("email")
	fmt.Println(email)

	productLink := c.Param("link")
	fmt.Println(productLink)

	collection := client.Database("octopusDB").Collection("user")
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"email": email},
		bson.M{"$pull": bson.M{"watchlist": bson.M{"link": productLink}}})

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please try again!"})
	} else if result != nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Removed from watchlist!"})

	}
}
