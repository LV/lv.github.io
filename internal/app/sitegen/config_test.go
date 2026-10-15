package sitegen

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSite(t *testing.T) {
	path := filepath.Join(t.TempDir(), "site.json")
	config := `{
		"title": "Title",
		"subtitle": "Subtitle",
		"author": "Author",
		"feed": {"title": "Feed Title"}
	}`
	if err := os.WriteFile(path, []byte(config), 0o644); err != nil {
		t.Fatal(err)
	}
	s, err := LoadSite(path)
	if err != nil {
		t.Fatal(err)
	}
	want := Site{
		Title:    "Title",
		Subtitle: "Subtitle",
		Author:   "Author",
		Feed:     Feed{Title: "Feed Title"},
	}
	if s != want {
		t.Errorf("LoadSite() = %+v, want %+v", s, want)
	}
}

// TestSiteConfig checks that configs/site.json sets every field, without
// testing what they're set to.
func TestSiteConfig(t *testing.T) {
	s, err := LoadSite("../../../configs/site.json")
	if err != nil {
		t.Fatal(err)
	}
	fields := map[string]string{
		"title":      s.Title,
		"subtitle":   s.Subtitle,
		"author":     s.Author,
		"feed.title": s.Feed.Title,
	}
	for name, value := range fields {
		if value == "" {
			t.Errorf("configs/site.json: %s is missing or empty", name)
		}
	}
	for _, key := range siteKeys(t) {
		if _, ok := fields[key]; !ok {
			t.Errorf("Site has %s, but TestSiteConfig doesn't check it; add it to fields", key)
		}
	}
}

// siteKeys returns every key in Site's JSON form, with nested keys joined by
// dots, e.g. "feed.title". A field tagged omitempty would be left out.
func siteKeys(t *testing.T) []string {
	t.Helper()
	b, err := json.Marshal(Site{})
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	var keys []string
	var walk func(prefix string, m map[string]any)
	walk = func(prefix string, m map[string]any) {
		for k, v := range m {
			if sub, ok := v.(map[string]any); ok {
				walk(prefix+k+".", sub)
			} else {
				keys = append(keys, prefix+k)
			}
		}
	}
	walk("", m)
	return keys
}

func TestLoadSiteUnknownField(t *testing.T) {
	path := filepath.Join(t.TempDir(), "site.json")
	if err := os.WriteFile(path, []byte(`{"titel": "LV"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadSite(path); err == nil {
		t.Error("LoadSite() with an unknown field: got nil error")
	}
}
