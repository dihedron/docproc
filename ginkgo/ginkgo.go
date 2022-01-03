package ginkgo

import (
	"fmt"
	"io"
	"path"
	"text/template"

	"github.com/dihedron/ginkgo/ginkgo/funcs"
)

type Engine struct {
	input     map[string]interface{}
	output    io.Writer
	templates []string
	main      string
	functions template.FuncMap
}

// Option is the type for functional options.
type Option func(*Engine)

// New creates a new Engine, applying all the provided functional options.
func New(options ...Option) *Engine {
	p := &Engine{
		input:     map[string]interface{}{},
		templates: []string{},
		functions: template.FuncMap{},
	}
	for _, option := range options {
		option(p)
	}
	return p
}

// Close closes the output writer, if it implements the method.
func (e *Engine) Close() error {
	if e.output != nil {
		if c, ok := e.output.(io.WriteCloser); ok {
			return c.Close()
		}
	}
	return nil
}

// WithInput allows to specify the Engine input map.
func WithInput(input map[string]interface{}) Option {
	return func(e *Engine) {
		if input != nil {
			e.input = input
		}
	}
}

// WithOutput allows to specify the Engine output stream.
func WithOutput(output io.Writer) Option {
	return func(e *Engine) {
		if output != nil {
			e.output = output
		}
	}
}

// WithTemplate allows to specify a template files.
func WithTemplate(main bool, template string) Option {
	return func(e *Engine) {
		if main {
			e.main = path.Base(template)
		}
		e.templates = append(e.templates, template)
	}
}

func WithDefaultFunctions() Option {
	return func(e *Engine) {
		e.functions["include"] = funcs.Include
	}
}

func WithFunction(name string, fx interface{}) Option {
	return func(e *Engine) {
		e.functions[name] = fx
	}
}

func (e *Engine) Execute() error {

	// parse the templates
	templates, err := template.New(e.main).Funcs(e.functions).ParseFiles(e.templates...)
	if err != nil {
		return fmt.Errorf("error parsing template files %v: %w", e.templates, err)
	}

	// execute the template
	if err := templates.ExecuteTemplate(e.output, e.main /*map[string]interface{}(*cmd.Input)*/, e.input); err != nil {
		return fmt.Errorf("error applying data to template: %w", err)
	}
	return nil
}
