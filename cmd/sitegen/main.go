package main

import (
	"log"
	"os"
)

const (
	dirPerm     = 0o755 // owner: read, write, open; everyone else: read, open
	outDir      = "_site"
	templateDir = "web/template"
)

func main() {
	if err := os.RemoveAll(outDir); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(outDir, dirPerm); err != nil {
		log.Fatal(err)
	}
	if err := buildHomepage(); err != nil {
		log.Fatal(err)
	}
}

// buildHomepage renders the homepage template into the output directory.
func buildHomepage() error {
}
