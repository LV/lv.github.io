package sitegen

// page pairs a template under site.DirPage with the data it renders with.
type page struct {
	path string
	data pageData
}

// pages lists every page the build renders. It is temporary: once pages are
// discovered from web/content, each page's data will live in its own file.
var pages = []page{
	{
		path: "index.html",
		data: pageData{
			Description: "Creating to Understand",
			CurrentNav:  "home",
		},
	},
	{
		path: "writings/index.html",
		data: pageData{
			Title:       "Writings",
			Description: "Thoughts and Experiments",
			CurrentNav:  "writings",
		},
	},
}
