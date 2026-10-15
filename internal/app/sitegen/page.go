// Package sitegen builds the site from web/content and web/template into dist/.
package sitegen

// Site holds the settings every page shares, loaded from configs/site.json.
type Site struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Author   string `json:"author"`
}

// Page is one HTML file in dist/, built from one Markdown file in web/content/.
type Page struct {
	Path   string // output path under dist/, e.g. "index.html", "blog/index.html"
	Source string // Markdown file under web/content/, e.g. "index.md"
}
