package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func search(c *gin.Context) {
	var sp SearchPost

	if err := c.BindJSON(&sp); err != nil {
		return
	}
	search_data, err := json.Marshal(&sp)

	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("http://localhost:8000", "application/json",
		bytes.NewBuffer(search_data))

	if err != nil {
		log.Fatal(err)
	}
	var sr []SearchResults

	json.NewDecoder(resp.Body).Decode(&sr)

	c.IndentedJSON(http.StatusCreated, sr)

}
