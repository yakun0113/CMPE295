package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func scrape_product_with_python() {

	cmd := exec.Command("/usr/local/bin/python", "/Users/brian80433/Desktop/pythonScraping/scraper.py")
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

func main() {
	start := time.Now()
	product := "beef"
	latitude := "37.3380242"
	longitude := "-121.9187964"
	info := product + "," + latitude + "," + longitude
	pass_product_name_and_location(info)
	scrape_product_with_python()
	duration := time.Since(start)
	fmt.Println(duration)
}
