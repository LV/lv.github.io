package sitegen

import (
	"errors"
	"io/fs"
	"path"
	"strings"
)

// Discover returns a Page for every Markdown file in content, in alphabetical
// order. Each x/y.md becomes x/y.html. The root must have an index.md.
func Discover(content fs.FS) ([]Page, error) {
	var pages []Page
	err := fs.WalkDir(content, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || path.Ext(p) != ".md" {
			return nil
		}
		pages = append(pages, Page{Path: strings.TrimSuffix(p, ".md") + ".html", Source: p})
		return nil
	})
	if err != nil {
		return nil, err
	}
	if _, err := fs.Stat(content, "index.md"); err != nil {
		return nil, errors.New("index.md is required")
	}
	return pages, nil
}
