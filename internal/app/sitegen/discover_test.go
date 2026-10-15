package sitegen

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestDiscover(t *testing.T) {
	content := fstest.MapFS{
		"index.md":         {},
		"blog/index.md":    {},
		"blog/first.md":    {},
		"images/photo.jpg": {},
		"notes.txt":        {},
	}
	pages, err := Discover(content)
	if err != nil {
		t.Fatal(err)
	}
	want := []Page{
		{Path: "blog/first.html", Source: "blog/first.md"},
		{Path: "blog/index.html", Source: "blog/index.md"},
		{Path: "index.html", Source: "index.md"},
	}
	if !reflect.DeepEqual(pages, want) {
		t.Errorf("Discover() = %+v, want %+v", pages, want)
	}
}

func TestDiscoverMissingIndex(t *testing.T) {
	content := fstest.MapFS{"blog/index.md": {}}
	if _, err := Discover(content); err == nil {
		t.Error("Discover() without index.md: got nil error")
	}
}
