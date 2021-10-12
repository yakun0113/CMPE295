package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func scrape_product_with_python() {

	cmd := exec.Command("/usr/local/bin/python", "/Users/brian80433/Desktop/CMPE295-main/backend/scraper.py")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))

}

func pass_product_name_and_location(info string) {

	f, err := os.Create("product_name.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(info)

	if err2 != nil {
		log.Fatal(err2)
	}

}

type searchPost struct {
	ItemName  string `json:"itemName"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var sp searchPost
	decoder.Decode(&sp)
	item := sp.ItemName
	latitude := sp.Latitude
	longitude := sp.Longitude

	info := item + "," + latitude + "," + longitude
	pass_product_name_and_location(info)
	scrape_product_with_python()

	data, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer data.Close()
	jdata, err := ioutil.ReadAll(data)
	w.Write(jdata)
}

func main() {

	fs := http.FileServer(http.Dir("../frontend/dist"))
	http.Handle("/", fs)
	http.HandleFunc("/search", searchHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServeTLS(":8080", "localhost.crt", "localhost.key", nil); err != nil {
		log.Fatal(err)
	}

}
