package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

func main() {
	options := Options{}

	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		fmt.Fprintf(os.Stderr, "error parsing command line: %v\n", err)
		os.Exit(1)
	}
	err := options.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error executing command: %v\n", err)
		os.Exit(1)
	}
}
