package ginkgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/dihedron/ginkgo/ginkgo/funcs"
	"gopkg.in/yaml.v3"
)

type Engine struct {
	Input *Input `short:"i" long:"input" description:"The input data, either as an inline JSON value or as a @file (in JSON or YAML format)." otional:"yes" env:"GINKGO_INPUT"`
	// Input     string   `short:"i" long:"input" description:"The path to the input file." optional:"yes" env:"GINKGO_INPUT"`
	// Format    string   `short:"f" long:"format" description:"The format of the input data, if provided via STDIN." choice:"yaml" choice:"json" default:"json" optional:"yes" env:"GINKGO_FORMAT"`
	Template  string   `short:"m" long:"main" description:"The name of the main template file." required:"yes" env:"GINKGO_TEMPLATE"`
	Output    string   `short:"o" long:"output" description:"The path to the output file." optional:"yes" env:"GINKGO_OUTPUT"`
	Templates []string `short:"t" long:"template" description:"The paths of all the templates and subtemplates on disk." required:"yes"`
}

func (cmd *Engine) Execute() error {
	var output *os.File

	// load all the templates
	templates, err := template.New(cmd.Template).Funcs(FuncMap).ParseFiles(cmd.Templates...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing template files %v: %v\n", cmd.Templates, err)
		return err
	}
	// if the input map is nil, then the input data is
	// provided via STDIN, and that's where we take it
	if cmd.Input == nil {
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("error reading input data from STDIN: %w", err)
		}
		m := Input(make(map[string]interface{}))
		cmd.Input = &m
		if strings.HasPrefix(strings.TrimLeft(string(input), " \n\r"), "---") {
			if err = yaml.Unmarshal(input, &cmd.Input); err != nil {
				return fmt.Errorf("error unmarshalling YAML input: %w", err)
			}
		} else if strings.HasPrefix(strings.TrimLeft(string(input), " \n\r"), "{") {
			if err = json.Unmarshal(input, &cmd.Input); err != nil {
				return fmt.Errorf("error unmarshalling JSON input: %w", err)
			}
		} else {
			return fmt.Errorf("unrecognisable input format on STDIN")
		}
	}
	// prepare output stream
	if cmd.Output != "" {
		path := filepath.Dir(cmd.Output)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "error creating output directory %s: %v\v", path, err)
			return err
		}
		output, err = os.Create(cmd.Output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating output file %s: %v\v", cmd.Output, err)
			return err
		}
	} else {
		output = os.Stdout
	}
	// execute the template
	if err := templates.ExecuteTemplate(output, cmd.Template, map[string]interface{}(*cmd.Input)); err != nil {
		fmt.Fprintf(os.Stderr, "error applying variables to template: %v\n", err)
		return err
	}

	return nil
}

// This FuncMap is used to register the custom functions.
var FuncMap = template.FuncMap{
	"include": funcs.Include,
}
