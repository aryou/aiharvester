package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/singlehopllc/aiharvester/beater"
)

func main() {
	err := beat.Run("aiharvester", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
