package main

import (
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const (
	permDir     = 0o755 // owner: read, write, open; everyone else: read, open
	permFile    = 0o644
	dirOut      = "_site"
	dirTemplate = "web/template"
	dirStatic   = "web/static"
)

func main() {
	if err := os.RemoveAll(dirOut); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(dirOut, permDir); err != nil {
		log.Fatal(err)
	}
	if err := buildHomepage(); err != nil {
		log.Fatal(err)
	}
	if err := writeStatic(); err != nil {
		log.Fatal(err)
	}
}

// buildHomepage renders the homepage template into the output directory.
func buildHomepage() error {
	tmpl, err := template.ParseFiles(filepath.Join(dirTemplate, "index.html"))
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(dirOut, "index.html"))
	if err != nil {
		return err
	}

	if err := tmpl.Execute(f, nil); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

// writeStatic writes every file under dirStatic into dirOut unchanged,
// keeping the same relative paths.
func writeStatic() error {
	return filepath.WalkDir(dirStatic, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(dirStatic, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dirOut, rel)

		if d.IsDir() {
			return os.MkdirAll(target, permDir)
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, permFile)
	})
}
