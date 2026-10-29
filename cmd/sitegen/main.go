package main

import (
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/LV/lv.github.io/internal/site"
)

const (
	permDir  = 0o755 // owner: read, write, open; everyone else: read, open
	permFile = 0o644 // owner: read, write; everyone else: read
)

func main() {
	if err := os.RemoveAll(site.DirOut); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(site.DirOut, permDir); err != nil {
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
	tmpl, err := template.ParseFiles(filepath.Join(site.DirTemplate, "index.html"))
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(site.DirOut, "index.html"))
	if err != nil {
		return err
	}

	if err := tmpl.Execute(f, nil); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

// writeStatic writes every file under site.DirStatic into site.DirOut unchanged,
// keeping the same relative paths.
func writeStatic() error {
	return filepath.WalkDir(site.DirStatic, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(site.DirStatic, path)
		if err != nil {
			return err
		}
		target := filepath.Join(site.DirOut, rel)

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
