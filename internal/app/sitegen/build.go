// Package sitegen builds the static site: it renders every page into the
// layout and copies the static files, writing the result to an output folder.
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

// Build reads templates and static files from in, clears the folder out, and
// regenerates the whole site into it. Every template sees now as the build
// time.
func Build(in fs.FS, out string, now time.Time) error {
	if err := os.RemoveAll(out); err != nil {
		return err
	}
	if err := os.MkdirAll(out, permDir); err != nil {
		return err
	}
	layout, err := parseBaseLayout(in, templateFuncs(now))
	if err != nil {
		return err
	}
	for _, p := range pages {
		if err := renderPage(layout, in, out, p.path, p.data); err != nil {
			return err
		}
	}
	return writeStatic(in, out)
}

// writeStatic writes every file under site.DirStatic within in into out
// unchanged, keeping the same relative paths.
func writeStatic(in fs.FS, out string) error {
	static, err := fs.Sub(in, site.DirStatic)
	if err != nil {
		return err
	}

	// Paths from static are already relative to site.DirStatic.
	return fs.WalkDir(static, ".", func(rel string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		target := filepath.Join(out, filepath.FromSlash(rel))

		if d.IsDir() {
			return os.MkdirAll(target, permDir)
		}

		data, err := fs.ReadFile(static, rel)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, permFile)
	})
}
