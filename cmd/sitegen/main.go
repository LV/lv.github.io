package main

import (
	"log"
	"os"
	"time"

	"github.com/LV/lv.github.io/internal/app/sitegen"
	"github.com/LV/lv.github.io/internal/pkg/site"
)

func main() {
	// UTC, so local and CI builds agree on the date.
	if err := sitegen.Build(os.DirFS("."), site.DirOut, time.Now().UTC()); err != nil {
		log.Fatal(err)
	}
}
