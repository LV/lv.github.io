// Package sitegen builds the static site: it renders every page into the
// layout and copies the static files, writing the result to site.DirOut.
package sitegen

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/LV/lv.github.io/internal/pkg/site"
)

const (
	permDir  = 0o755 // owner: read, write, open; everyone else: read, open
	permFile = 0o644 // owner: read, write; everyone else: read
)

// Build clears site.DirOut and regenerates the whole site into it. Every
// template sees now as the build time.
func Build(now time.Time) error {
	if err := os.RemoveAll(site.DirOut); err != nil {
		return err
	}
	if err := os.MkdirAll(site.DirOut, permDir); err != nil {
		return err
	}
	layout, err := parseBaseLayout(templateFuncs(now))
	if err != nil {
		return err
	}
	for _, p := range pages {
		if err := renderPage(layout, p.path, p.data); err != nil {
			return err
		}
	}
	return writeStatic()
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
