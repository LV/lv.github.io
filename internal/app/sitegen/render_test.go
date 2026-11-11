package sitegen

import (
	"os"
	"path"
	"path/filepath"
	"testing"
	"testing/fstest"
	"time"

	"github.com/LV/lv.github.io/internal/pkg/site"
)

func TestRenderPage(t *testing.T) {
	fsys := fstest.MapFS{
		path.Join(site.DirLayout, "base.html"): {Data: []byte(`<title>{{.Title}}</title>{{template "content" .}} {{year}}`)},
		path.Join(site.DirPage, "a.html"):      {Data: []byte(`{{define "content"}}A{{end}}`)},
		path.Join(site.DirPage, "sub/b.html"):  {Data: []byte(`{{define "content"}}B{{end}}`)},
	}
	now := time.Date(2030, time.June, 1, 0, 0, 0, 0, time.UTC)

	layout, err := parseBaseLayout(fsys, templateFuncs(now))
	if err != nil {
		t.Fatal(err)
	}

	outDir := t.TempDir()
	tests := []struct {
		rel  string
		data pageData
		want string
	}{
		{"a.html", pageData{Title: "One"}, "<title>One</title>A 2030"},
		{"sub/b.html", pageData{Title: "Two"}, "<title>Two</title>B 2030"},
	}

	// Render every page from the same layout first, then check them, so a page
	// leaking into the layout would show up in a later page's output.
	for _, tt := range tests {
		if err := renderPage(layout, fsys, outDir, tt.rel, tt.data); err != nil {
			t.Fatalf("renderPage(%q): %v", tt.rel, err)
		}
	}
	for _, tt := range tests {
		if got := readFile(t, filepath.Join(outDir, tt.rel)); got != tt.want {
			t.Errorf("%s = %q, want %q", tt.rel, got, tt.want)
		}
	}
}

func readFile(t *testing.T, name string) string {
	t.Helper()
	data, err := os.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}
