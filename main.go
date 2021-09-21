package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
)

func main() {
	options := Options{}
	log.Println("entering application")

	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		log.Printf("error parsing command line: %v\n", err)
		os.Exit(1)
	}
	err := options.Execute()
	if err != nil {
		log.Printf("error executing command: %v\n", err)
		os.Exit(1)
	}
}
