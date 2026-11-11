package sitegen

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/LV/lv.github.io/internal/pkg/site"
)

// pageData holds the values a page fills into the layout.
type pageData struct {
	Title       string // page name shown after "LV - " in <title>; empty for just "LV"
	Description string // <meta name="description">: the snippet search engines show under the title
	CurrentNav  string
}

// parseBaseLayout returns the base layout parsed with funcs registered. It is
// the master that every page is cloned from (see renderPage), so funcs reach
// every page too. Never execute it directly: Clone fails once a template has
// been executed.
func parseBaseLayout(funcs template.FuncMap) (*template.Template, error) {
	path := filepath.Join(site.DirLayout, "base.html")

	// Funcs must be registered before parsing. Execute runs the template named
	// in New, so that name must be the layout's file name.
	return template.New(filepath.Base(path)).Funcs(funcs).ParseFiles(path)
}

// renderPage parses the page template at site.DirPage/rel into a clone of
// layout, executes it with data, and writes the result to site.DirOut/rel,
// creating any folders the output path needs. layout itself is left unchanged.
func renderPage(layout *template.Template, rel string, data pageData) error {
	pathTmpl := filepath.Join(site.DirPage, rel)
	pathOut := filepath.Join(site.DirOut, rel)

	// Each page gets its own clone of the layout, so the "content" it defines
	// can't replace another page's.
	tmpl, err := layout.Clone()
	if err != nil {
		return err
	}
	if _, err := tmpl.ParseFiles(pathTmpl); err != nil {
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
