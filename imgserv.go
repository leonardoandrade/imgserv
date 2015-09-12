package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func status(w http.ResponseWriter, r *http.Request) {
	log.Print("'" + r.URL.Path + "' called")
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
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

func main() {
	http.HandleFunc("/status", status)
	http.HandleFunc("/u", upload)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
