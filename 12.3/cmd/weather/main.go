package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"weather"
)

func main() {
	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		log.Fatal("Please set the environment variable OPENWEATHERMAP_API_KEY.")
	}
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s LOCATION\n\nExample: %[1]s London,UK", os.Args[0])
	}
	location := os.Args[1]
	URL := weather.FormatURL(weather.BaseURL, location, key)
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal("unexpected response status ", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	conditions, err := weather.ParseResponse(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conditions)
}
