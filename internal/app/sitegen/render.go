package sitegen

import (
	"html/template"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/LV/lv.github.io/internal/pkg/site"
)

// pageData holds the values a page fills into the layout.
type pageData struct {
	Title       string // page name shown after "LV - " in <title>; empty for just "LV"
	Description string // <meta name="description">: the snippet search engines show under the title
	CurrentNav  string
}

// parseBaseLayout returns the base layout from fsys, parsed with funcs
// registered. It is the master that every page is cloned from (see
// renderPage), so funcs reach every page too. Never execute it directly: Clone
// fails once a template has been executed.
func parseBaseLayout(fsys fs.FS, funcs template.FuncMap) (*template.Template, error) {
	layoutPath := path.Join(site.DirLayout, "base.html")

	// Funcs must be registered before parsing. Execute runs the template named
	// in New, so that name must be the layout's file name.
	return template.New(path.Base(layoutPath)).Funcs(funcs).ParseFS(fsys, layoutPath)
}

// renderPage parses the page template at site.DirPage/rel within fsys into a
// clone of layout, executes it with data, and writes the result to outDir/rel,
// creating any folders the output path needs. layout itself is left unchanged.
func renderPage(layout *template.Template, fsys fs.FS, outDir, rel string, data pageData) error {
	tmplPath := path.Join(site.DirPage, rel)
	outPath := filepath.Join(outDir, filepath.FromSlash(rel))

	// Each page gets its own clone of the layout, so the "content" it defines
	// can't replace another page's.
	tmpl, err := layout.Clone()
	if err != nil {
		return err
	}
	if _, err := tmpl.ParseFS(fsys, tmplPath); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(outPath), permDir); err != nil {
		return err
	}

	f, err := os.Create(outPath)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(f, data); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
