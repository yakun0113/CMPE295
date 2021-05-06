package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gocolly/colly"
)

type Coupon struct {
	Name      string `json:"name"`
	Location  string `json:"location"`
	Old_Price string `json:"old price"`
	Discount  string `json:"discount"`
}

func scraper(w http.ResponseWriter, r *http.Request) {
	var city string
	fmt.Scanln(&city)

	allCoupons := make([]Coupon, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("livingsocial.com", "www.livingsocial.com"),
	)

	collector.OnHTML("figure", func(element *colly.HTMLElement) {

		var name string
		var location string
		var old_price string
		var discount string

		name = element.ChildText(".grpn-dc-title")
		location = element.ChildText(".grpn-dc-loci")
		old_price = element.ChildText(".wh-dc-price-original")
		discount = element.ChildText(".wh-dc-discount")

		coupon := Coupon{
			Name:      name,
			Location:  location,
			Old_Price: old_price,
			Discount:  discount,
		}
		allCoupons = append(allCoupons, coupon)

	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://www.livingsocial.com/local/" + city)

	sort.SliceStable(allCoupons, func(i, j int) bool {
		return strings.TrimRight((allCoupons[i].Discount), "% discount_off") > strings.TrimRight((allCoupons[j].Discount), "% discount_off")
	})

	file, err := json.MarshalIndent(allCoupons, "", " ")
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
		http.ListenAndServe(":8000", nil)
	}
}
