package main

import (
	"fmt"
	"os"

	_ "github.com/dihedron/mason/autolog"
	"github.com/dihedron/mason/command"
	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"
)

func main() {
	defer zap.L().Sync()

	options := command.Commands{}

	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		zap.S().With(zap.Error(err)).Error("failure parsing command line")
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
