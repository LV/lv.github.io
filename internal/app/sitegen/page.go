// Package sitegen builds the site from web/content and web/template into dist/.
package sitegen

import "html/template"

// Site holds the defaults every page inherits (header title/description).
type Site struct {
	Title       string
	Description string
	Author      string
	Pages       []Page
}

// Page is one HTML file in dist/.
type Page struct {
	Path        string        // output path under dist/, e.g. "index.html", "blog/index.html"
	Title       string        // overrides Site.Title when non-empty
	Description string        // overrides Site.Description when non-empty
	Body        template.HTML // rendered content
}
