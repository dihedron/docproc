package main

import (
	"os"

	"github.com/dihedron/ginkgo/ginkgo"
	"github.com/jessevdk/go-flags"
)

func main() {
	ginkgo := ginkgo.Engine{}

	parser := flags.NewParser(&ginkgo, flags.Default)
	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
	err := ginkgo.Execute()
	if err != nil {
		os.Exit(1)
	}
}
