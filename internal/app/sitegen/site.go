package sitegen

// site is the whole site: its defaults and every page it builds.
var site = Site{
	Title:       "LV",
	Description: "Creating to Understand",
	Author:      "Luis Victoria",
	Pages: []Page{
		{Path: "index.html"},
	},
}
