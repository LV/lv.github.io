package main

import (
	"log"
	"time"

	"github.com/LV/lv.github.io/internal/app/sitegen"
)

func main() {
	// UTC, so local and CI builds agree on the date.
	if err := sitegen.Build(time.Now().UTC()); err != nil {
		log.Fatal(err)
	}
}
