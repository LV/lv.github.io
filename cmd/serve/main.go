package main

import (
	"log"
	"net/http"
)

const (
	addr   = "localhost:8080"
	dirOut = "_site"
)

func main() {
	handler := http.FileServer(http.Dir(dirOut))

	log.Printf("serving %s at http://%s", dirOut, addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
