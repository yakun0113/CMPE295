package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../../frontend/dist"))
	http.Handle("/", fs)
	http.HandleFunc("/search", search)
	http.HandleFunc("/signUp", signUp)
	http.HandleFunc("/signIn", signIn)

	//start server
	fmt.Printf("Starting server at port 8080\n")
	connectDB()
	if err := http.ListenAndServeTLS(":8080", "localhost.crt", "localhost.key", nil); err != nil {
		log.Fatal(err)
	}

}
