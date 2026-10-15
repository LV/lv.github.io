package sitegen

import (
	"encoding/json"
	"os"
)

// LoadSite reads the site settings from the JSON file at path.
// Unknown fields are an error, so a typo in the config fails the build.
func LoadSite(path string) (Site, error) {
	f, err := os.Open(path)
	if err != nil {
		return Site{}, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()

	var s Site
	if err := dec.Decode(&s); err != nil {
		return Site{}, err
	}
	return s, nil
}
