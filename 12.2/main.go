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
	if key == "" {
		log.Fatal("Please set the environment variable OPENWEATHERMAP_API_KEY.")
	}
	URL := fmt.Sprintf("%s/data/2.5/weather?q=London,UK&appid=%s", BaseURL, key)
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal("unexpected response status ", resp.Status)
	}
	io.Copy(os.Stdout, resp.Body)
}
