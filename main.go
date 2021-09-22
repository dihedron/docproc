package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

func main() {
	options := Options{}

	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
	err := options.Execute()
	if err != nil {
		os.Exit(1)
	}
}
