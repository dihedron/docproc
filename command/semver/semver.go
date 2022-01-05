package semver

// SemanticVersion is the set of root command groups.
type SemanticVersion struct {
	// Bump bumps the version (major, minor or revsion/patch) and prints it out.
	Bump Bump `command:"bump" alias:"b" description:"Bump the version in a semantic version."`
	// Check parses the version and prints it out.
	Check Check `command:"check" alias:"c" description:"Check the semantic version by parsing it."`
}
