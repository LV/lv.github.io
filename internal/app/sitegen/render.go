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

// renderPage renders the page template at site.DirPage/page inside the base
// layout, with funcs available to both, and writes the result to
// site.DirOut/page, creating any folders the output path needs.
func renderPage(page string, data pageData, funcs template.FuncMap) error {
	pathLayout := filepath.Join(site.DirLayout, "base.html")
	pathTmpl := filepath.Join(site.DirPage, page)
	pathOut := filepath.Join(site.DirOut, page)

	// Funcs must be registered before parsing. Execute runs the template named
	// in New, so that name must be the layout's file name.
	tmpl, err := template.New(filepath.Base(pathLayout)).Funcs(funcs).ParseFiles(pathLayout, pathTmpl)
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
