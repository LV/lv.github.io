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

// pageData holds the values a page fills into the layout.
type pageData struct {
	Title       string
	Description string
	Current     string // which nav link the header marks as current
}

func main() {
	if err := os.RemoveAll(site.DirOut); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(site.DirOut, permDir); err != nil {
		log.Fatal(err)
	}
	if err := renderPage("index.html", pageData{
		Title:       "LV",
		Description: "Creating to Understand",
		Current:     "home",
	}); err != nil {
		log.Fatal(err)
	}
	if err := renderPage("blog/index.html", pageData{
		Title:       "LV - Blog",
		Description: "Thoughts and Experiments",
		Current:     "blog",
	}); err != nil {
		log.Fatal(err)
	}
	if err := writeStatic(); err != nil {
		log.Fatal(err)
	}
}

// renderPage renders the page template at site.DirPage/page inside the base
// layout, and writes the result to site.DirOut/page, creating any folders
// the output path needs.
func renderPage(page string, data pageData) error {
	pathLayout := filepath.Join(site.DirLayout, "base.html")
	pathTmpl := filepath.Join(site.DirPage, page)
	pathOut := filepath.Join(site.DirOut, page)

	// The layout must come first: Execute runs the first file's template.
	tmpl, err := template.ParseFiles(pathLayout, pathTmpl)
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

	if err := tmpl.Execute(f, data); err != nil {
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
