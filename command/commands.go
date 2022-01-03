package command

import (
	"github.com/dihedron/mason/command/hydrate"
	"github.com/dihedron/mason/command/version"
)

// Commands is the set of root command groups.
type Commands struct {
	// Version prints the application version and exits.
	Version version.Version `command:"version" alias:"ver" alias:"v" description:"Print the command version and exit."`
	// Hydrate fills a template with data from an input structure in YAML or JSON format.
	Hydrate hydrate.Hydrate `command:"hydrate" alias:"hyd" alias:"h" description:"Hydrate a set of templates with input data."`
}
