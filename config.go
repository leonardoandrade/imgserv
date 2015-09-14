package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Directory      string   `json:"directory"`
	SupportedTypes []string `json:"supportedTypes"`
	SupportedSizes []string `json:"supportedSizes"`
}

func MakeDefaultConfig(directory string) Config {
	ret := Config{}
	ret.Directory = directory
	ret.SupportedTypes = []string{"png", "jpg"}
	ret.SupportedSizes = []string{"200x200", "300x300"}
	return ret
}

func MakeConfigFromFile(filePath string) (Config, error) {
	jsonContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(jsonContent, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c *Config) WriteToFile(filePath string) error {
	jsonBytes, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, jsonBytes, 0755)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) ToJson() ([]byte, error) {
	return json.MarshalIndent(c, "", "\t")
}
