package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Directory      string   `json:"directory"`
	SupportedTypes []string `json:"supportedTypes"`
}

func MakeDefaultConfig(directory string) Config {
	ret := Config{}
	ret.Directory = directory
	ret.SupportedTypes = []string{"png", "jpg"}
	return ret
}

func ConfigFromFile(filePath string) (Config, error) {
	//TODO
	return Config{}, nil
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
