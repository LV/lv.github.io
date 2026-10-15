package sitegen

import (
	"errors"
	"io/fs"
)

// RequiredInputs lists the files the generator reads, as paths from the repo
// root. The build can't run without them.
var RequiredInputs = []string{
	"configs/site.json",
	"web/content/index.md",
	"web/template/base.html",
}

// RequiredAssets lists the files the built site links to, as paths from the
// repo root. The generator doesn't read them, but they're copied into the
// output, and without them the site is broken.
var RequiredAssets = []string{
	"web/static/favicon.ico",
	"web/static/favicon-dark.ico",
	"web/static/index.css",
}

// CheckRequired returns an error naming every file in RequiredInputs and
// RequiredAssets that root is missing, or nil if they all exist. It runs
// before any other build work.
func CheckRequired(root fs.FS) error {
	var errs []error
	for _, name := range RequiredInputs {
		if _, err := fs.Stat(root, name); err != nil {
			errs = append(errs, errors.New("missing generator input: "+name))
		}
	}
	for _, name := range RequiredAssets {
		if _, err := fs.Stat(root, name); err != nil {
			errs = append(errs, errors.New("missing site asset: "+name))
		}
	}
	return errors.Join(errs...)
}
