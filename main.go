package main

import (
	"os"

	"github.com/elastic/beats/libbeat/cmd/instance"

	"github.com/Psiphon-Inc/cloudwatchlogsbeat/beater"
)

func main() {
	settings := instance.Settings{
		Name: "cloudwatchlogsbeat",
	}
	err := instance.Run(settings, beater.Creator)
	if err != nil {
		os.Exit(1)
	}
}
