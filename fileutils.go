package main

import (
	"math/rand"
)

var hex_chars = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E"}

type Metadata struct {
	original_filename string  `json:"originalFilename"`
	coordinates       *[2]int `json:"coordinates"`
}

func check_directory_emptyness(directory string) error {
	//TODO
	return nil
}

func random_guid() string {
	var ret [16]byte
	for i := 0; i < 32; i++ {
		ret[i] = hex_chars[rand.Intn(len(hex_chars))][0]
	}
	return string(ret[:])
}

func save_file(directory string, filename string, config Config) string {
	//TODO
	return ""
}
