package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	PORT int = 8888
)

func status(repository Repository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		statusJson, err := json.MarshalIndent(repository.Status(), "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(statusJson)
	}
}

func download(repository Repository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("guid: %s, width: %s", r.FormValue("guid"), r.FormValue("w"))

		imgBytes, err := repository.getImage(r.FormValue("guid"), r.FormValue("w"))
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Length", strconv.Itoa(len(imgBytes)))
		if _, err := w.Write(imgBytes); err != nil {
			log.Println("unable to write image.")
		}


	}
}

func upload(repository Repository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("'" + r.URL.Path + "' called")
		fmt.Printf("Request: %v\n", r)
		fn, header, err := r.FormFile("datafile")
		if err != nil {
			fmt.Printf("error %v\n",err)
		}
		if fn == nil || header == nil {
			fmt.Printf("request is not multipart")
			return
		}

		id, err := repository.saveImage(header.Filename, &fn)
		fmt.Printf("id: %s err: %s", id, err)
	}
}

func imgserv_usage() {
	fmt.Printf("Launch image server. The only parameter is the directory.\n")
	fmt.Printf("Usage:\n")
	fmt.Printf("imgserv <directory>")
}

func main() {
	if len(os.Args) != 2 {
		imgserv_usage()
		os.Exit(0)
	}
	repository := MakeRepository(os.Args[1])

	http.HandleFunc("/status", status(repository))
	http.HandleFunc("/d/", download(repository))
	http.HandleFunc("/u", upload(repository))

	fmt.Printf("Serving image directory '%s' on port '%d'\n", os.Args[1], PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}
