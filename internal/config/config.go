package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ServerAddress string `json:"server_address"`
	Target        string `json:"target"`
}

func LoadConfig(cfgPath string) (*Config, error) {
	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode config file: %v", err)
	}
	return &config, nil
}
