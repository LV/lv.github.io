package sitegen

import (
	"io/fs"
	"path"
	"strings"
)

// Discover returns a Page for every Markdown file in content, in alphabetical
// order. Each x/y.md becomes x/y.html. Callers run CheckRequired first, which
// ensures index.md exists.
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
	return pages, nil
}
