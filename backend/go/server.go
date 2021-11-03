package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx context.Context

//var collection *mongo.Collection

var dbUri = "mongodb+srv://Brian80433:cmpe295@octopus.oolyk.mongodb.net/octopusDB?retryWrites=true&w=majority"

func connectDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	//collection := client.Database("octopusDB").Collection("user")

}

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("../../frontend/dist", true)))
	router.POST("/handleSearch", search)
	router.POST("/handleSignUp", signUp)
	router.POST("/handleSignIn", signIn)

	router.GET("/handleWatchlist/:email", getItems)
	router.PUT("/handleWatchlist/:email", addItem)
	router.DELETE("/handleWatchlist/:email", deleteItem)
	/*
		authorized := router.Group("/")
		authorized.Use(AuthRequired)
		{
			authorized.GET("/handleAccount/:email")
			authorized.PUT("/handleAccount/:email")
			authorized.DELETE("/handleAccount/:email")

			authorized.GET("/handleWatchlist/:email", getItems)
			authorized.PUT("/handleWatchlist", addItem)
			authorized.DELETE("/handleWatchlist/:email", deleteItem)
		}
	*/
	connectDB()
	fmt.Printf("Starting server at port 8080\n")

	log.Fatal(router.RunTLS(":8080", "localhost.crt", "localhost.key"))
}
