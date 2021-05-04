package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func scraper(w http.ResponseWriter, r *http.Request) {
	var animal string
	fmt.Scanln(&animal)

	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factId, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("Could not get id")
		}

		factDesc := element.Text

		fact := Fact{
			ID:          factId,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://www.factretriever.com/" + animal + "-facts")

	sort.SliceStable(allFacts, func(i, j int) bool {
		return allFacts[i].ID < allFacts[j].ID
	})

	file, err := json.MarshalIndent(allFacts, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(file)
}

func main() {
	http.HandleFunc("/", scraper)
	// check if port variable has been set
	if os.Getenv("PORT") == "" {
		http.ListenAndServe(":9000", nil)
	}
}
