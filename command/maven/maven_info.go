package maven

import (
	"html/template"
	"os"
)

type Info struct {
	Command
}

func (cmd *Info) Execute(args []string) error {
	// cmd.Project.GroupId.Text

	t := template.Must(template.New("").Parse(`Group ID    : {{ .GroupId.Text }}
Artifact ID : {{ .ArtifactId.Text }}
Version     : {{ .Version.Text }}
`))
	t.Execute(os.Stdout, cmd.Project)
	return nil
}
