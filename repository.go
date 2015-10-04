package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"github.com/gographics/imagick"
	"fmt"
	"os"
	"io"
	//"image"
)

var hex_chars = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E"}

type RepositoryStatus struct {
	Directory string `json:"directory"`
	Config    Config `json:"config"`
	CacheSize int    `json:"cache_size"`
}

type Repository struct {
	directory string
	config    Config
	cache     map[string]bool //simplistic cache
}

type FileMetadata struct {
	OriginalName string  `json:"original_name"`
	Location     *[2]int `json:"location"`
}

func MakeRepository(directory string) Repository {
	config, err := MakeConfigFromFile(directory + "/config.json")
	if err != nil {
		panic(err)
	}
	cache := make(map[string]bool)
	return Repository{directory, config, cache}
}

func (r *Repository) getImage(guid, width string) ([]byte, error) {

	filePath := r.directory+"/"+string(guid[0])+"/"+string(guid[1])+"/"+guid+"/original.png"
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *Repository) makeGuid() string {
	var ret []byte = make([]byte, 32)
	for i := 0; i < 32; i++ {
		ret[i] = hex_chars[rand.Intn(len(hex_chars)-1)][0]
	}
	return string(ret[:])
}

func (r *Repository) Status() RepositoryStatus {
	return RepositoryStatus{r.directory, r.config, len(r.cache)}
}

type Metadata struct {
	original_filename string  `json:"originalFilename"`
	original_size []int `json:"originalSize"`
	coordinates       *[2]int `json:"coordinates"`
}

func check_directory_emptyness(directory string) error {
	//TODO
	return nil
}

func (r *Repository) processAndSaveImage(imageRoot string, blob []byte) (error) {
	fmt.Printf("xxxxx%d\n",len(blob))

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	//err := mw.ReadImageBlob(blob)
	err := mw.ReadImage(imageRoot+"/original.png")
	if err != nil {
		fmt.Printf("image not valid")
	}
	w := mw.GetImageWidth()
	h := mw.GetImageHeight()

	fmt.Printf("width: %d, height: %d\n",w, h)
	fmt.Printf("format: %s\n", mw.GetFormat())

	defer mw.Destroy()

	return nil
}

func (r *Repository) saveImage(filename string, content *multipart.File) (string, error) {
	randName := r.makeGuid()

	defer (*content).Close()
	imageRoot := r.directory + "/" + string(randName[0]) + "/" + string(randName[1]) + "/" + randName
	os.MkdirAll(imageRoot, 0755)
	f, _ := os.Create(imageRoot + "/original.png")
	io.Copy(f, (*content))
	f.Close()

	var blob = make([]byte, 100000)
	(*content).Read(blob)
	r.processAndSaveImage(imageRoot, blob)
	metadata := FileMetadata{filename, &[2]int{0, 0}}
	jsonBytes, err := json.MarshalIndent(metadata, "", "\t")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(imageRoot+"/meta.json", jsonBytes, 0755)
	if err != nil {
		return "", err
	}

	return randName, nil
}
