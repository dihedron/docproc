package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Options struct {
	Input    string `short:"i" long:"input" description:"The path to the input file." optional:"yes" env:"GINKGO_INPUT"`
	Format   string `short:"f" long:"format" description:"The format of the input data, if provided via STDIN." choice:"yaml" choice:"json" default:"json" optional:"yes" env:"GINKGO_FORMAT"`
	Template string `short:"t" long:"template" description:"The path to the template file." required:"yes" env:"GINKGO_TEMPLATE"`
	Output   string `short:"o" long:"output" description:"The path to the output file." optional:"yes" env:"GINKGO_OUTPUT"`
}

func (opts *Options) Execute() error {

	var input []byte
	var output *os.File
	var err error

	log.Println("entering command")

	// read the template
	templ, err := template.ParseFiles(opts.Template)
	if err != nil {
		return err
	}

	// read input
	if opts.Input != "" {
		log.Printf("reading from input: %v\n", opts.Input)
		input, err = os.ReadFile(opts.Input)
	} else {
		log.Println("reading from standard input")
		input, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		fmt.Printf("error reading input: %v\n", err)
		return err
	}

	// prepare output stream
	if opts.Output != "" {
		output, err = os.Create(opts.Output)
	} else {
		output = os.Stdout
	}
	if err != nil {
		return err
	}

	// find the input file format
	var format string
	if opts.Input != "" {
		switch strings.ToLower(filepath.Ext(opts.Input)) {
		case "yaml", "yml", ".yml", ".yaml":
			log.Println("extension: YAML")
			format = "yaml"
		case "json", ".json":
			log.Println("extension: JSON")
			format = "json"
		default:
			log.Fatalf("unsupported input format: %v", filepath.Ext(opts.Input))
		}
	} else {
		format = opts.Format
	}

	// read in input and unmarshal it
	dynamic := make(map[string]interface{})
	switch format {
	case "yaml":
		log.Println("extension: YAML")
		if err = yaml.Unmarshal(input, &dynamic); err != nil {
			return err
		}
	case "json":
		log.Println("extension: JSON")
		if err = json.Unmarshal(input, &dynamic); err != nil {
			return err
		}
	default:
		log.Println("extension", strings.ToLower(filepath.Ext(opts.Input)))
	}

	if err = templ.Execute(output, dynamic); err != nil {
		return err
	}

	return nil
}
