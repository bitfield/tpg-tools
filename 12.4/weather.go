package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Conditions struct {
	Summary string
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
}

func ParseResponse(data []byte) (Conditions, error) {
	var resp OWMResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Conditions{}, fmt.Errorf("invalid API response %s: %v", data, err)
	}
	if len(resp.Weather) < 1 {
		return Conditions{}, fmt.Errorf("invalid API response %s: want at least one Weather element", data)
	}
	conditions := Conditions{
		Summary: resp.Weather[0].Main,
	}
	return conditions, nil
}

type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		BaseURL:    "https://api.openweathermap.org",
		HTTPClient: http.DefaultClient,
	}
}

func (c Client) FormatURL(location string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", c.BaseURL, location, c.APIKey)
}

func (c *Client) GetWeather(location string) (Conditions, error) {
	URL := c.FormatURL(location)
	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		return Conditions{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Conditions{}, fmt.Errorf("unexpected response status %q", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Conditions{}, err
	}
	conditions, err := ParseResponse(data)
	if err != nil {
		return Conditions{}, fmt.Errorf("invalid API response %s: %v", data, err)
	}
	return conditions, nil
}

func RunCLI() {
	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "Please set the environment variable OPENWEATHERMAP_API_KEY.")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s LOCATION\n\nExample: %[1]s London,UK\n", os.Args[0])
		os.Exit(1)
	}
	location := os.Args[1]
	c := NewClient(key)
	conditions, err := c.GetWeather(location)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conditions)
}
