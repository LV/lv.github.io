// Package sitegen builds the site from web/content and web/template into dist/.
package sitegen

// Site holds the defaults every page inherits (header title/description).
type Site struct {
	Title       string
	Description string
	Author      string
	Pages       []Page
}

// Page is one HTML file in dist/, built from one Markdown file in web/content/.
type Page struct {
	Path        string // output path under dist/, e.g. "index.html", "blog/index.html"
	Source      string // Markdown file under web/content/, e.g. "index.md"
	Title       string // overrides Site.Title when non-empty
	Description string // overrides Site.Description when non-empty
}
