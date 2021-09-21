package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

	// read the template
	templ, err := template.ParseFiles(opts.Template)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing template file %s: %v\n", opts.Template, err)
		return err
	}

	// read input
	if opts.Input != "" {
		input, err = os.ReadFile(opts.Input)
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		return err
	}

	// prepare output stream
	if opts.Output != "" {
		path := filepath.Dir(opts.Output)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "error creating output directory %s: %v\v", path, err)
			return err
		}
		output, err = os.Create(opts.Output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating output file %s: %v\v", opts.Output, err)
			return err
		}
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
			format = "yaml"
		case "json", ".json":
			format = "json"
		default:
			fmt.Fprintf(os.Stderr, "unsupported input format: %s\n", filepath.Ext(opts.Input))
		}
	} else {
		format = opts.Format
	}

	// read in input and unmarshal it
	dynamic := make(map[string]interface{})
	switch format {
	case "yaml":
		if err = yaml.Unmarshal(input, &dynamic); err != nil {
			fmt.Fprintf(os.Stderr, "error unmarshalling YAML input: %v\n", err)
			return err
		}
	case "json":
		if err = json.Unmarshal(input, &dynamic); err != nil {
			fmt.Fprintf(os.Stderr, "error unmarshalling JSON input: %v\n", err)
			return err
		}
	default:
		fmt.Fprintf(os.Stderr, "unsupported input format: %s\n", format)
		return errors.New("unsupported input format")
	}

	if err = templ.Execute(output, dynamic); err != nil {
		fmt.Fprintf(os.Stderr, "error applying variables to template: %v\n", err)
		return err
	}

	return nil
}
