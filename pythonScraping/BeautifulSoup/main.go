package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func scrape_product_with_python() {

	cmd := exec.Command("/usr/local/bin/python", "/Users/brian80433/Desktop/pythonScraping/BeautifulSoup/scraper.py")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))

}

func pass_product_name(product string) {

	f, err := os.Create("product_name.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(product)

	if err2 != nil {
		log.Fatal(err2)
	}

}

func main() {

	product := "toilet-paper"
	pass_product_name(product)
	scrape_product_with_python()

}
