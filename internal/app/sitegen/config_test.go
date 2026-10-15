package sitegen

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSite(t *testing.T) {
	s, err := LoadSite("../../../configs/site.json")
	if err != nil {
		t.Fatal(err)
	}
	want := Site{Title: "LV", Subtitle: "Creating to Understand", Author: "Luis Victoria"}
	if s != want {
		t.Errorf("LoadSite() = %+v, want %+v", s, want)
	}
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
