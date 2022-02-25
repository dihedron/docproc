package manual

import (
	_ "embed"
	"html/template"
	"os"

	"github.com/dihedron/mason/build"
	"github.com/dihedron/mason/command/base"
	"github.com/dihedron/mason/command/version"
)

type Manual struct {
	base.Command
}

//go:embed template.tpl
var manual string

// Execute is the real implementation of the Manual command.
func (cmd *Manual) Execute(args []string) error {

	helper := &version.DetailedInfo{
		Name:        build.Name,
		Description: build.Description,
		Version:     version.VersionInfo{Major: build.VersionMajor, Minor: build.VersionMinor, Patch: build.VersionPatch},
	}

	tmpl := template.Must(template.New("manualTemplate").Parse(manual))
	err := tmpl.Execute(os.Stdout, helper)
	if err != nil {
		panic(err)
	}
	return nil
}
