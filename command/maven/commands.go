package maven

import "github.com/dihedron/mason/command/base"

type Command struct {
	base.Command
	// Input is the optional input value, either as a file (prepend with @) or as an inline value (start with --- with)
	Project *Project `short:"p" long:"pom" description:"The POM to process, either as an inline XML value or as a @file." otional:"yes" env:"MASON_INPUT"`
}

// Commands is the set of root command groups.
type Maven struct {
	// Version prints the application version and exits.
	Info Info `command:"info" alias:"i" description:"Print general info about the given POM."`
}
