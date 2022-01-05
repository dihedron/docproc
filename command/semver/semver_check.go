package semver

import (
	"github.com/dihedron/mason/command/base"
)

type Check struct {
	base.Command
}

func (cmd *Check) Execute(args []string) error {
	// command Bump with no increment indication
	cmd2 := &Bump{
		Command: base.Command{
			Automation: cmd.Automation,
			Parameters: cmd.Parameters,
		},
		Major:    false,
		Minor:    false,
		Patch:    false,
		Revision: false,
	}
	return cmd2.Execute(args)
}
