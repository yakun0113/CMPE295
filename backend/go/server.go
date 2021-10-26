package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("../../frontend/dist", true)))
	router.POST("/search", search)
	router.POST("/signUp", signUp)
	router.POST("/signIn", signIn)

	authorized := router.Group("/")
	authorized.Use(AuthRequired)
	authorized.GET("/userProfile")
	authorized.POST("/userProfile")
	authorized.PUT("/userProfile")
	authorized.DELETE("/userProfile")

	authorized.GET("/watchlist")
	authorized.POST("/watchlist")
	authorized.PUT("/watchlist")
	authorized.DELETE("/watchlist")

	connectDB()

	router.RunTLS(":8080", "localhost.crt", "localhost.key")

	/*
			fs := http.FileServer(http.Dir("../../frontend/dist"))
		http.Handle("/", fs)
		http.HandleFunc("/search", search)
		http.HandleFunc("/signUp", signUp)
		http.HandleFunc("/signIn", signIn)

		fmt.Printf("Starting server at port 8080\n")
		connectDB()
		if err := http.ListenAndServeTLS(":8080", "localhost.crt", "localhost.key", nil); err != nil {
			log.Fatal(err)
		}

	*/
}
