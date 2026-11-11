package main

import (
	"testing"
	"time"
)

func TestYearsBetween(t *testing.T) {
	from := mustParse(t, "1998-12-04T15:12:00-04:00")

	tests := []struct {
		name string
		to   string
		want int
	}{
		{"before this year's birthday", "2026-11-11T00:00:00Z", 27},
		{"one minute before", "2026-12-04T19:11:00Z", 27},
		{"exact moment", "2026-12-04T19:12:00Z", 28},
		{"new year in UTC, still Dec 31 in AST", "2027-01-01T02:00:00Z", 28},
		{"birth moment", "1998-12-04T19:12:00Z", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := yearsBetween(from, mustParse(t, tt.to)); got != tt.want {
				t.Errorf("yearsBetween(%s, %s) = %d, want %d", from, tt.to, got, tt.want)
			}
		})
	}
}

func mustParse(t *testing.T, ts string) time.Time {
	t.Helper()
	parsed, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		t.Fatal(err)
	}
	return parsed
}
