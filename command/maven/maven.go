package maven

// Maven is the set of root command groups.
type Maven struct {
	// Info prints the application version and exits.
	Info Info `command:"info" alias:"i" description:"Print general info about the given POM."`
}
