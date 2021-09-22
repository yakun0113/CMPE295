package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

/*
product := "beef"
latitude := "37.3380242"
longitude := "-121.9187964"
*/

func scrape_product_with_python() {

	cmd := exec.Command("/usr/local/bin/python", "scraper.py")
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

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	item := r.FormValue("item")
	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")

	fmt.Fprintf(w, "Item = %s\n", item)
	fmt.Fprintf(w, "Latitude = %s\n", latitude)
	fmt.Fprintf(w, "Longitude = %s\n", longitude)

	info := item + "," + latitude + "," + longitude
	pass_product_name_and_location(info)
	scrape_product_with_python()
	// Open our jsonFile
	data, err := os.Open("products.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened products.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer data.Close()
	jdata, err := ioutil.ReadAll(data)

	w.Write(jdata)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServeTLS(":8080", "localhost.crt", "localhost.key", nil); err != nil {
		log.Fatal(err)
	}
}
