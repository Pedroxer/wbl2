package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config содержит порт, на котором работает web-server
type Config struct {
	Port string `json:"port"`
}

// ParseConfig парсит конфигурационный файл и возвращает *Config
func ParseConfig(configPath string) (*Config, error) {
	fileBody, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(fileBody, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
