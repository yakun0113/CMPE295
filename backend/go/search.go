package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type searchPost struct {
	ItemName  string `json:"itemName"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type searchResults struct {
	Store    string     `json:"store"`
	Location location   `json:"location"`
	Products []products `json:"products"`
}

type location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type products struct {
	ID     int32  `json:"id"`
	Image  string `json:"image"`
	Name   string `json:"name"`
	Price  string `json:"price"`
	Rating string `json:"rating"`
	Link   string `json:"link"`
}

func search(c *gin.Context) {
	var sp searchPost

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
	var sr []searchResults

	json.NewDecoder(resp.Body).Decode(&sr)

	c.IndentedJSON(http.StatusCreated, sr)

}
