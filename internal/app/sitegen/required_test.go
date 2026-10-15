package sitegen

import (
	"os"
	"strings"
	"testing"
	"testing/fstest"
)

func TestCheckRequired(t *testing.T) {
	root := fstest.MapFS{}
	for _, name := range Required {
		root[name] = &fstest.MapFile{}
	}
	if err := CheckRequired(root); err != nil {
		t.Errorf("CheckRequired() with every file: %v", err)
	}
}

func TestCheckRequiredMissing(t *testing.T) {
	root := fstest.MapFS{}
	for _, name := range Required {
		root[name] = &fstest.MapFile{}
	}
	delete(root, "web/content/index.md")
	delete(root, "web/static/favicon.ico")

	err := CheckRequired(root)
	if err == nil {
		t.Fatal("CheckRequired() with missing files: got nil error")
	}
	for _, name := range []string{"web/content/index.md", "web/static/favicon.ico"} {
		if !strings.Contains(err.Error(), name) {
			t.Errorf("CheckRequired() error %q doesn't name %s", err, name)
		}
	}
}

func TestCheckRequiredRepo(t *testing.T) {
	if err := CheckRequired(os.DirFS("../../..")); err != nil {
		t.Errorf("CheckRequired() on the repo: %v", err)
	}
}
