package sitegen

import (
	"errors"
	"io/fs"
)

// Required lists the files the build can't run without, as paths from the
// repo root. CheckRequired verifies them before any other build work.
var Required = []string{
	"configs/site.json",
	"web/content/index.md",
	"web/template/base.html",
	"web/static/favicon.ico",
	"web/static/favicon-dark.ico",
	"web/static/index.css",
}

// CheckRequired returns an error naming every file in Required that root is
// missing, or nil if they all exist.
func CheckRequired(root fs.FS) error {
	var errs []error
	for _, name := range Required {
		if _, err := fs.Stat(root, name); err != nil {
			errs = append(errs, errors.New("missing required file: "+name))
		}
	}
	return errors.Join(errs...)
}
