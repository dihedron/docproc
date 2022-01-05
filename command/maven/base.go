package maven

import "github.com/dihedron/mason/command/base"

// Command is the base command
type Command struct {
	base.Command
	// Input is the optional input value, either as a file (prepend with @) or as an inline value (start with --- with)
	POM *POM `short:"p" long:"pom" description:"The POM to process, either as an inline XML value or as a @file." otional:"yes" env:"MASON_INPUT"`
}
