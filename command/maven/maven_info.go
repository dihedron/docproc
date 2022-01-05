package maven

import (
	_ "embed"
	"html/template"
	"os"
)

type Info struct {
	Command
}

//go:embed templates/simple.tpl
var simple string

func (cmd *Info) Execute(args []string) error {
	t := template.Must(template.New("").Parse(simple))
	t.Execute(os.Stdout, cmd.POM)
	return nil
}
