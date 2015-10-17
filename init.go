package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Initialize a directory with sub-directory tree and default configuration file. Must be empty.\n")
	fmt.Printf("Usage:\n")
	fmt.Printf("init <directory>")
}

/*
* Create empty directory with subfolder structure and default configuration.
 */
func main() {

	if len(os.Args) != 2 {
		usage()
		os.Exit(0)
	}
	directory := os.Args[1]
	fmt.Printf("Creating directory structure in '%s' for image server.\n", directory)

	if nil != check_directory_emptyness(directory) {
		//TODO
		os.Exit(0)
	}
	for i := 0; i < len(hex_chars); i++ {
		for j := 0; j < len(hex_chars); j++ {
			path := directory + "/" + hex_chars[i] + "/" + hex_chars[j]
			err := os.MkdirAll(path, 0755)
			if err != nil {
				fmt.Print("Error creating %s: %s", directory, err)
				os.Exit(0)
			}
		}
	}

	configFilePath := directory + "/config.json"
	config := MakeDefaultConfig(directory)
	config.WriteToFile(configFilePath)
	fmt.Printf("Directories and '" + configFilePath + "' config file created with success. \n")
}
