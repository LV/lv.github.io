package sitegen

import (
	"errors"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"testing"
	"testing/fstest"
	"time"

	"github.com/LV/lv.github.io/internal/pkg/site"
)

func TestBuild(t *testing.T) {
	in := fstest.MapFS{
		path.Join(site.DirLayout, "base.html"):    {Data: []byte(`{{template "content" .}}`)},
		path.Join(site.DirStatic, "styles/a.css"): {Data: []byte("body {}")},
	}
	// One template per entry in pages, whose content is its own path.
	for _, p := range pages {
		in[path.Join(site.DirPage, p.path)] = &fstest.MapFile{
			Data: []byte(`{{define "content"}}` + p.path + `{{end}}`),
		}
	}

	out := t.TempDir()
	stale := filepath.Join(out, "stale.html")
	if err := os.WriteFile(stale, []byte("old"), permFile); err != nil {
		t.Fatal(err)
	}

	if err := Build(in, out, time.Date(2030, time.June, 1, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Fatal(err)
	}

	for _, p := range pages {
		if got := readFile(t, filepath.Join(out, p.path)); got != p.path {
			t.Errorf("%s = %q, want %q", p.path, got, p.path)
		}
	}
	if got := readFile(t, filepath.Join(out, "styles", "a.css")); got != "body {}" {
		t.Errorf("styles/a.css = %q, want %q", got, "body {}")
	}
	if _, err := os.Stat(stale); !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("stale file still exists after Build (stat err: %v)", err)
	}
}
