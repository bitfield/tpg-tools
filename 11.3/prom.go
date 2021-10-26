package prom

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Global GlobalConfig
}

type GlobalConfig struct {
	ScrapeInterval     time.Duration     `yaml:"scrape_interval"`
	EvaluationInterval time.Duration     `yaml:"evaluation_interval"`
	ScrapeTimeout      time.Duration     `yaml:"scrape_timeout"`
	ExternalLabels     map[string]string `yaml:"external_labels"`
}

func ConfigFromYAML(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()
	config := Config{
		GlobalConfig{
			ScrapeTimeout: 10 * time.Second,
		},
	}
	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
