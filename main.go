package main

import (
	"os"

	"github.com/Psiphon-Inc/cloudwatchlogsbeat/beater"

	"github.com/elastic/beats/libbeat/beat"
)

func main() {
	err := beat.Run("cloudwatchlogsbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
