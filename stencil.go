package main

import (
	"os"

	log "github.com/Sirupsen/logrus"

	"bitbucket.org/lgslab/stencil/opts"
	"bitbucket.org/lgslab/stencil/render"
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
		render.Render(opts.Parse_argument(arg))
	}

}
