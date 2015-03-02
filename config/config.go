package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	ApiKey string
	Goals  []*Goal
}

type Goal struct {
	Type       string
	TimeWindow string
	GoalTime   string
}

func NewConfig(path string) (config *Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New(fmt.Sprintf("Config file does not exist: %s", path))
		} else if os.IsPermission(err) {
			return nil, errors.New(fmt.Sprintf("Config file is unreadable: %s", path))
		}
		return nil, err
	}

	decoder := json.NewDecoder(file)
	config = &Config{}

	if err = decoder.Decode(&config); err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to parse config file: %s", path))
	}

	if err = validateConfig(config); err != nil {
		return nil, err
	}

	return config, err
}

func validateConfig(config *Config) error {
	errorMessage := ""

	if config.ApiKey == "" {
		errorMessage += "API Key "
	}

	if len(config.Goals) == 0 {
		errorMessage += "Goals "
	}

	if errorMessage != "" {
		return errors.New(fmt.Sprintf("Missing configuration values: %v", errorMessage))
	}

	return nil
}
