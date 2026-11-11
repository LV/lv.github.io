package sitegen

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/LV/lv.github.io/internal/site"
)

// pageData holds the values a page fills into the layout.
type pageData struct {
	Title       string // page name shown after "LV - " in <title>; empty for just "LV"
	Description string // <meta name="description">: the snippet search engines show under the title
	CurrentNav  string
}

// parseLayout parses the base layout once, with funcs available to it and to
// every page later parsed into a copy of it.
func parseLayout(funcs template.FuncMap) (*template.Template, error) {
	path := filepath.Join(site.DirLayout, "base.html")

	// Funcs must be registered before parsing. Execute runs the template named
	// in New, so that name must be the layout's file name.
	return template.New(filepath.Base(path)).Funcs(funcs).ParseFiles(path)
}

// renderPage renders the page template at site.DirPage/page into a copy of
// layout, and writes the result to site.DirOut/page, creating any folders the
// output path needs.
func renderPage(layout *template.Template, page string, data pageData) error {
	pathTmpl := filepath.Join(site.DirPage, page)
	pathOut := filepath.Join(site.DirOut, page)

	// Each page gets its own copy of the layout, so the "content" it defines
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
