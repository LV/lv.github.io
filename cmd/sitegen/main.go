package main

import (
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const (
	dirPerm     = 0o755 // owner: read, write, open; everyone else: read, open
	filePerm    = 0o644
	outDir      = "_site"
	templateDir = "web/template"
	staticDir   = "web/static"
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
	if err := writeStatic(); err != nil {
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

// writeStatic writes every file under staticDir into outDir unchanged,
// keeping the same relative paths.
func writeStatic() error {
	return filepath.WalkDir(staticDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(staticDir, path)
		if err != nil {
			return err
		}
		target := filepath.Join(outDir, rel)

		if d.IsDir() {
			return os.MkdirAll(target, dirPerm)
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, filePerm)
	})
}
