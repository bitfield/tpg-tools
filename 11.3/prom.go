package prom

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Label struct {
	Name, Value string
}

type GlobalConfig struct {
	ScrapeInterval     time.Duration     `yaml:"scrape_interval"`
	EvaluationInterval time.Duration     `yaml:"evaluation_interval"`
	ExternalLabels     map[string]string `yaml:"external_labels"`
}

type Config struct {
	Global GlobalConfig
}

func ConfigFromYAML(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()
	var config Config
	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
