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
	fsys := fstest.MapFS{
		path.Join(site.DirLayout, "base.html"):    {Data: []byte(`{{template "content" .}}`)},
		path.Join(site.DirStatic, "styles/a.css"): {Data: []byte("body {}")},
	}
	// One template per entry in pages, whose content is its own path.
	for _, p := range pages {
		fsys[path.Join(site.DirPage, p.path)] = &fstest.MapFile{
			Data: []byte(`{{define "content"}}` + p.path + `{{end}}`),
		}
	}

	outDir := t.TempDir()
	stale := filepath.Join(outDir, "stale.html")
	if err := os.WriteFile(stale, []byte("old"), permFile); err != nil {
		t.Fatal(err)
	}

	if err := Build(fsys, outDir, time.Date(2030, time.June, 1, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Fatal(err)
	}

	for _, p := range pages {
		if got := readFile(t, filepath.Join(outDir, p.path)); got != p.path {
			t.Errorf("%s = %q, want %q", p.path, got, p.path)
		}
	}
	if got := readFile(t, filepath.Join(outDir, "styles", "a.css")); got != "body {}" {
		t.Errorf("styles/a.css = %q, want %q", got, "body {}")
	}
	if _, err := os.Stat(stale); !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("stale file still exists after Build (stat err: %v)", err)
	}
}
