package main

import (
	"log"
	"net/http"

	"github.com/LV/lv.github.io/internal/site"
)

const (
	addr = "localhost:8080"
)

func main() {
	handler := http.FileServer(http.Dir(site.DirOut))

	log.Printf("serving %s at http://%s", site.DirOut, addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
