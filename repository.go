package main

import (
	"math/rand"
	"mime/multipart"
	"io"
	"os"
)

var hex_chars = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E"}

type RepositoryStatus struct {
	Directory      string   `json:"directory"`
	Config    Config `json:"config"`
	CacheSize int    `json:"cache_size"`
}

type Repository struct {
	directory string
	config Config
	cache  map[string]bool //simplistic cache
}


func MakeRepository(directory string) (Repository) {
	cache := make(map[string]bool)
	config, err := MakeConfigFromFile(directory + "/config.json")
	if err != nil {
		panic(err)
	}
	return Repository{directory, config, cache}
}

func (r * Repository) makeGuid() (string){
	var ret []byte = make([]byte, 32)
	for i := 0; i < 32; i++ {
		ret[i] = hex_chars[rand.Intn(len(hex_chars)-1)][0]
	}
	return string(ret[:])
}

func (r * Repository) Status() (RepositoryStatus) {
	return RepositoryStatus{r.directory, r.config, len(r.cache)}
}

type Metadata struct {
	original_filename string  `json:"originalFilename"`
	coordinates       *[2]int `json:"coordinates"`
}

func check_directory_emptyness(directory string) error {
	//TODO
	return nil
}


func (r * Repository) saveImage(filename string, content *multipart.File) (string, error) {
	randName := r.makeGuid()


	defer (*content).Close()

	imageRoot := r.directory + "/" + string(randName[0]) + "/" + string(randName[1]) + "/"+ randName
	os.MkdirAll(imageRoot, 0755)
	f, _ := os.Create(imageRoot + "/original.png")
	defer f.Close()

	io.Copy(f, (*content))

	return randName, nil
}
