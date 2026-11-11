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

// Build reads templates and static files from fsys, clears the folder outDir,
// and regenerates the whole site into it. Every template sees now as the build
// time.
func Build(fsys fs.FS, outDir string, now time.Time) error {
	if err := os.RemoveAll(outDir); err != nil {
		return err
	}
	if err := os.MkdirAll(outDir, permDir); err != nil {
		return err
	}
	layout, err := parseBaseLayout(fsys, templateFuncs(now))
	if err != nil {
		return err
	}
	for _, p := range pages {
		if err := renderPage(layout, fsys, outDir, p.path, p.data); err != nil {
			return err
		}
	}
	return writeStatic(fsys, outDir)
}

// writeStatic writes every file under site.DirStatic within fsys into outDir
// unchanged, keeping the same relative paths.
func writeStatic(fsys fs.FS, outDir string) error {
	static, err := fs.Sub(fsys, site.DirStatic)
	if err != nil {
		return err
	}

	// Paths from static are already relative to site.DirStatic.
	return fs.WalkDir(static, ".", func(rel string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		target := filepath.Join(outDir, filepath.FromSlash(rel))

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
