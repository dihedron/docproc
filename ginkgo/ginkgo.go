package ginkgo

import (
	"bufio"
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

type Engine struct {
	Input     string   `short:"i" long:"input" description:"The path to the input file." optional:"yes" env:"GINKGO_INPUT"`
	Format    string   `short:"f" long:"format" description:"The format of the input data, if provided via STDIN." choice:"yaml" choice:"json" default:"json" optional:"yes" env:"GINKGO_FORMAT"`
	Template  string   `short:"m" long:"main" description:"The name of the main template file." required:"yes" env:"GINKGO_TEMPLATE"`
	Output    string   `short:"o" long:"output" description:"The path to the output file." optional:"yes" env:"GINKGO_OUTPUT"`
	Templates []string `short:"t" long:"template" description:"The paths of all the templates and subtemplates on disk." required:"yes"`
}

const FormatYAML = "yaml"
const FormatJSON = "json"

func (cmd *Engine) Execute() error {

	var input []byte
	var output *os.File
	var format string
	var err error

	// load all the templates
	templates, err := template.New(cmd.Template).Funcs(FuncMap).ParseFiles(cmd.Templates...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing template files %v: %v\n", cmd.Templates, err)
		return err
	}

	// read input
	switch cmd.Input {
	case "": // read from STDIN
		input, err = ioutil.ReadAll(os.Stdin)
	case "---":
		input = []byte(cmd.Input)
		format = FormatYAML
	case "{}":
		input = []byte(cmd.Input)
		format = FormatJSON
	default: // file on disk
		input, err = os.ReadFile(cmd.Input)

	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		return err
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
	if err != nil {
		return err
	}

	// find the input file format, unless it's already
	// prepopulated by empty input
	if cmd.Input != "" {
		if format == "" { // not prepopulated by fake input ('{}' or '---')
			switch strings.ToLower(filepath.Ext(cmd.Input)) {
			case "yaml", "yml", ".yml", ".yaml":
				format = FormatYAML
			case "json", ".json":
				format = FormatJSON
			default:
				fmt.Fprintf(os.Stderr, "unsupported input format: %s\n", filepath.Ext(cmd.Input))
			}
		}
	} else {
		format = cmd.Format
	}

	// read in input and unmarshal it
	dynamic := make(map[string]interface{})
	switch format {
	case FormatYAML:
		if err = yaml.Unmarshal(input, &dynamic); err != nil {
			fmt.Fprintf(os.Stderr, "error unmarshalling YAML input: %v\n", err)
			return err
		}
	case FormatJSON:
		if err = json.Unmarshal(input, &dynamic); err != nil {
			fmt.Fprintf(os.Stderr, "error unmarshalling JSON input: %v\n", err)
			return err
		}
	default:
		fmt.Fprintf(os.Stderr, "unsupported input format: %s\n", format)
		return errors.New("unsupported input format")
	}

	//if err = templates.Execute(output, dynamic); err != nil {
	if err = templates.ExecuteTemplate(output, cmd.Template, dynamic); err != nil {
		fmt.Fprintf(os.Stderr, "error applying variables to template: %v\n", err)
		return err
	}

	return nil
}

// This FuncMap is used to register the custom functions.
var FuncMap = template.FuncMap{
	"include": Include,
}

func DumpArgs(args ...interface{}) (string, error) {
	result := ""
	if args != nil {
		for i, arg := range args {
			result += fmt.Sprintf("%d => '%v' (%T)\n", i, arg, arg)
		}
		fmt.Println(result)
		return result, nil
	} else {
		return "<empty>", nil
	}
}

// Include is the function that implements inclusion of
// subfiles with an optional padding; when used without
// padding it is roughly equivalent to "template"; padding
// provides a way to prepend a constant string to each line
// in the output. The usage is as follows:
// {{ include <template> [<pipeline>] [<padding>] }}
func Include(args ...interface{}) (string, error) {
	var (
		file    string
		padding string
		dynamic map[string]interface{}
	)

	if args == nil {
		return "", errors.New("include: at least the template path must be specified")
	}
	var pipelineFound bool
	for i, arg := range args {
		var ok bool

		if i == 0 {
			if file, ok = arg.(string); !ok {
				return "", errors.New("include: the first argument (template) must be of type string")
			}
		} else if i == 1 {
			if dynamic, ok = arg.(map[string]interface{}); !ok {
				if padding, ok = arg.(string); !ok {
					return "", errors.New("include: the second argument must either the pipeline or the padding")
				}
			} else {
				pipelineFound = true
			}
		} else if i == 2 {
			if !pipelineFound {
				return "", errors.New("include: the pipeline has not been provided")
			}
			if padding, ok = arg.(string); !ok {
				return "", errors.New("include: the third argument (padding) must be of type string")
			}
		}
	}

	// load the template
	t, err := template.ParseFiles(file)
	if err != nil {
		return "", err
	}

	var buffer strings.Builder
	if err = t.Execute(&buffer, dynamic); err != nil {
		return "", err
	}

	text := buffer.String()

	// apply padding only if necessary
	if padding != "" {
		var output strings.Builder
		scanner := bufio.NewScanner(strings.NewReader(text))
		for scanner.Scan() {
			output.WriteString(padding)
			output.WriteString(scanner.Text())
			output.WriteString("\n")
		}
		if scanner.Err() != nil {
			return "", err
		}
		return output.String(), nil
	}

	return text, nil
}
