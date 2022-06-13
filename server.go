package main

import (
	"log"
	"net/http"
)

const (
	port = ":4201"
	dir  = "./"
)

func main() {
	log.Printf("listening on %q...", port)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(dir))))
}
