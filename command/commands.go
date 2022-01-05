package command

import (
	"github.com/dihedron/mason/command/hydrate"
	"github.com/dihedron/mason/command/log"
	"github.com/dihedron/mason/command/maven"
	"github.com/dihedron/mason/command/semver"
	"github.com/dihedron/mason/command/version"
)

// Commands is the set of root command groups.
type Commands struct {
	// Version prints the application version and exits.
	Version version.Version `command:"version" alias:"ver" alias:"v" description:"Print the command version and exit."`
	// Hydrate fills a template with data from an input structure in YAML or JSON format.
	Hydrate hydrate.Hydrate `command:"hydrate" alias:"hyd" alias:"h" description:"Hydrate a set of templates using input data."`
	// Maven collects all Maven (Java) related commands.
	Maven maven.Maven `command:"maven" alias:"mvn" alias:"m" description:"Manipulate POM files."`
	// SemanticVersion collects all semantic versioning related commands.
	SemanticVersion semver.SemanticVersion `command:"semver" alias:"sv" alias:"s" description:"Manipulate semantic versions."`

	Log log.Log `command:"log" alias:"l" description:"Log messages to the console."`
}
