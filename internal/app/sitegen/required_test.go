package sitegen

import (
	"strings"
	"testing"
	"testing/fstest"
)

// allRequired returns a fake repo that has every file in RequiredInputs and
// RequiredAssets.
func allRequired() fstest.MapFS {
	root := fstest.MapFS{}
	for _, name := range append(RequiredInputs, RequiredAssets...) {
		root[name] = &fstest.MapFile{}
	}
	return root
}

func TestCheckRequired(t *testing.T) {
	if err := CheckRequired(allRequired()); err != nil {
		t.Errorf("CheckRequired() with every file: %v", err)
	}
}

func TestCheckRequiredMissing(t *testing.T) {
	root := allRequired()
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
