package semver

import (
	"github.com/dihedron/mason/command/base"
)

type QrCode struct {
	base.Command
	Input  Input  `short:"i" long:"output" description:"The input data to process" optional:"yes"`
	Output string `short:"o" long:"output" description:"The output file (empty for STDOUT)" optional:"yes"`
	Size   int    `short:"s" long:"size" description:"The size of the output image in pixels" optional:"yes" default:"256"`
	// TextArt specifies whether the application should output as text to the console.
	TextArt bool `short:"t" long:"text-art" description:"Whether the output should go to the console" optional:"yes"`
	// Invert specifies whether foreground (black) and background (white) colour should be inverted.
	Invert bool `short:"i" long:"invert" description:"Invert foreground and background colour" optional:"yes"`
	// Borderless specifies whether the QRCode should be created without a border.
	Borderless bool `short:"b" long:"no-border" description:"Create the QRCode without a border" optional:"yes"`
}

func (cmd *QrCode) Execute(args []string) error {
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
