package semver

// SemanticVersion is the set of root command groups.
type SemanticVersion struct {
	// Version prints the application version and exits.
	Bump Bump `command:"bump" alias:"b" description:"Bump the version in a semantic version."`
}
