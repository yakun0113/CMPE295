package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"octopus/models"

	"github.com/gin-gonic/gin"
)

func Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sp models.SearchPost

		if err := c.BindJSON(&sp); err != nil {
			return
		}
		search_data, err := json.Marshal(&sp)

		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.Post("http://localhost:5000", "application/json",
			bytes.NewBuffer(search_data))

		if err != nil {
			log.Fatal(err)
		}
		var sr []models.SearchResults

		json.NewDecoder(resp.Body).Decode(&sr)

		c.IndentedJSON(http.StatusCreated, sr)

	}
}
