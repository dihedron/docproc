package command

import (
	"git.gld-1.cloud.bankit.it/shared/utils/cloudctl/command/credits"
	"git.gld-1.cloud.bankit.it/shared/utils/cloudctl/command/json"
	"git.gld-1.cloud.bankit.it/shared/utils/cloudctl/command/plugin"
	"git.gld-1.cloud.bankit.it/shared/utils/cloudctl/command/repository"
	"git.gld-1.cloud.bankit.it/shared/utils/cloudctl/sdk/command/version"
)

// Commands is the set of root command groups.
type Commands struct {
	// Plugin is the namespace of all plugin related commands.
	Plugin plugin.Plugin `command:"plugin" alias:"plg" alias:"p" description:"Manage locally installed plugins."`
	// Repository provides the functionalities related to repository handling.
	Repository repository.Repository `command:"repository" alias:"repo" alias:"r" description:"Manage the the plugin repository."`
	// Version prints the application version and exits.
	Version version.Version `command:"version" alias:"ver" alias:"v" description:"Print the command version and exit."`
	// Credits prints the application credits and exits.
	Credits credits.Credits `command:"credits" alias:"cred" alias:"c" description:"Print credits and exit." hidden:"yes"`
	// Json parses a sample JSON file to a custom struct and prints it out to console.
	Json json.Json `command:"json" alias:"j" hidden:"yes" description:"Parse a JSON file out to console."`
}
