package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const BaseURL = "https://api.openweathermap.org"

func main() {
	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	URL := fmt.Sprintf("%s/data/2.5/weather?q=London,UK&appid=%s", BaseURL, key)
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
