package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	PORT int = 8888
)

type Server struct {
	Config Config
	Cache  map[string]bool
}

type ServerStatus struct {
	Config    Config `json:"config"`
	CacheSize int    `json:"cache_size"`
}

func MakeServer(directory string) Server {
	cache := make(map[string]bool)
	config, err := MakeConfigFromFile(directory + "/config.json")
	if err != nil {
		panic(err)
	}
	return Server{config, cache}
}

func status(server Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		statusJson, err := json.MarshalIndent(server.GetStatus(), "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(statusJson)
	}
}

func (s *Server) GetStatus() ServerStatus {
	return ServerStatus{s.Config, len(s.Cache)}
}

func upload(w http.ResponseWriter, r *http.Request) {
	log.Print("'" + r.URL.Path + "' called")
	fmt.Printf("Request: %v\n", r)
	fn, header, _ := r.FormFile("datafile")
	defer fn.Close()

	f, _ := os.Create("/tmp/" + header.Filename)
	defer f.Close()

	io.Copy(f, fn)
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func config(w http.ResponseWriter, r *http.Request) {
	log.Print("'" + r.URL.Path + "' called")
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func usage() {
	fmt.Printf("Launch image server. The only parameter is the directory.\n")
	fmt.Printf("Usage:\n")
	fmt.Printf("imgserv <directory>")
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(0)
	}
	server := MakeServer(os.Args[1])

	http.HandleFunc("/status", status(server))
	fmt.Printf("Serving image directory '%s' on port '%d'\n", os.Args[1], PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))

}
