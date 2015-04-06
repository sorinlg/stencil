package main

import (
	"os"

	log "github.com/Sirupsen/logrus"

	"./opts"
	"./template"
)

// setup log
func init() {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Rendering templates...")

	args := os.Args[1:]

	for _, arg := range args {
		template.Render(opts.Parse_argument(arg))
	}
}
