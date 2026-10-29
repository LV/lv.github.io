package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
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
	tmpl, err := template.ParseFiles(filepath.Join(templateDir, "index.html"))
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(outDir, "index.html"))
	if err != nil {
		return err
	}

	if err := tmpl.Execute(f, nil); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
