package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Getenv("AUTH_URL")

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}
