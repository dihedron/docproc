package hydrate

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/dihedron/mason/command/base"
	"github.com/dihedron/mason/command/hydrate/formatting"
	"github.com/dihedron/mason/unmarshal"
	"gopkg.in/yaml.v3"
)

type Hydrate struct {
	base.Command
	Input     *Input   `short:"i" long:"input" description:"The input data, either as an inline JSON value or as a @file (in JSON or YAML format)." otional:"yes" env:"MASON_INPUT"`
	Templates []string `short:"t" long:"template" description:"The paths of all the templates and subtemplates on disk." required:"yes"`
	Output    string   `short:"o" long:"output" description:"The path to the output file." optional:"yes" env:"MASON_OUTPUT"`
}

type Input map[string]interface{}

func (i *Input) UnmarshalFlag(value string) error {
	return unmarshal.FromFlag(value, i)
}

func (cmd *Hydrate) Execute(args []string) error {
	var err error

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

	// prepare the output stream
	var output io.Writer
	if cmd.Output != "" {
		path := filepath.Dir(cmd.Output)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("error creating output directory %s: %w", path, err)
		}
		output, err = os.Create(cmd.Output)
		if err != nil {
			return fmt.Errorf("error creating output file %s: %w", cmd.Output, err)
		}
	} else {
		output = os.Stdout
	}

	// populate the functions map
	functions := template.FuncMap{}
	for k, v := range formatting.FuncMap() {
		functions[k] = v
	}
	for k, v := range sprig.FuncMap() {
		functions[k] = v
	}

	// parse the templates
	main := path.Base(cmd.Templates[0])
	templates, err := template.New(main).Funcs(functions).ParseFiles(cmd.Templates...)
	if err != nil {
		return fmt.Errorf("error parsing template files %v: %w", cmd.Templates, err)
	}

	// execute the template
	if err := templates.ExecuteTemplate(output, main, map[string]interface{}(*cmd.Input)); err != nil {
		return fmt.Errorf("error applying data to template: %w", err)
	}
	return nil
}
