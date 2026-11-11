package sitegen

import (
	"html/template"
	"time"
)

// templateFuncs returns the functions every template can call. They all read
// the same now, so every page in one build agrees on the date.
func templateFuncs(now time.Time) template.FuncMap {
	return template.FuncMap{
		"year": func() int { return now.Year() },

		// yearsSince returns how many full years have passed since ts (RFC 3339).
		"yearsSince": func(ts string) (int, error) {
			t, err := time.Parse(time.RFC3339, ts)
			if err != nil {
				return 0, err
			}
			return yearsBetween(t, now), nil
		},
	}
}

// yearsBetween counts the full years between from and to. It uses from's time
// zone for the calendar, so each year ends at the exact anniversary moment.
func yearsBetween(from, to time.Time) int {
	to = to.In(from.Location())
	years := to.Year() - from.Year()
	if to.Before(from.AddDate(years, 0, 0)) {
		years-- // this year's anniversary hasn't happened yet
	}
	return years
}
