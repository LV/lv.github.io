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
	if err := renderPage("index.html"); err != nil {
		log.Fatal(err)
	}
	if err := renderPage("blog/index.html"); err != nil {
		log.Fatal(err)
	}
	if err := writeStatic(); err != nil {
		log.Fatal(err)
	}
}

// renderPage renders the template at site.DirTemplate/page into
// site.DirOut/page, creating any folders the output path needs.
func renderPage(page string) error {
	pathTmpl := filepath.Join(site.DirTemplate, page)
	pathOut := filepath.Join(site.DirOut, page)

	tmpl, err := template.ParseFiles(pathTmpl)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(pathOut), permDir); err != nil {
		return err
	}

	f, err := os.Create(pathOut)
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
