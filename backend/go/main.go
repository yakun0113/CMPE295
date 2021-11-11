package main

import (
	"log"
	"octopus/routes"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(static.Serve("/", static.LocalFile("../../frontend/dist", true)))

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	//router.Run(":" + port)
	log.Fatal(router.RunTLS(":"+port, "localhost.crt", "localhost.key"))

}
