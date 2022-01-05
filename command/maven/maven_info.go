package maven

import (
	_ "embed"
	"fmt"
	"html/template"
	"os"
)

type Info struct {
	Command
	GroupId    bool `short:"g" long:"group-id" description:"Print the group id." optional:"yes"`
	ArtifactId bool `short:"a" long:"artifact-id" description:"Print the artifact id." optional:"yes"`
	Version    bool `short:"v" long:"version" description:"Print the version." optional:"yes"`
}

//go:embed templates/simple.tpl
var simple string

func (cmd *Info) Execute(args []string) error {

	if (cmd.GroupId && cmd.ArtifactId) || (cmd.GroupId && cmd.Version) || (cmd.ArtifactId && cmd.Version) {
		return fmt.Errorf("only one of Group ID (value: %t), Artifact ID (value: %t) and Version (value: %t) can be requested at once", cmd.GroupId, cmd.ArtifactId, cmd.Version)
	}

	if cmd.GroupId {
		fmt.Printf("%s\n", cmd.POM.GroupId.Text)
	} else if cmd.ArtifactId {
		fmt.Printf("%s\n", cmd.POM.ArtifactId.Text)
	} else if cmd.Version {
		fmt.Printf("%s\n", cmd.POM.Version.Text)
	} else {
		t := template.Must(template.New("").Parse(simple))
		t.Execute(os.Stdout, cmd.POM)
	}
	return nil
}
