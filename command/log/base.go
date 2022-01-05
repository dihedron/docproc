package log

// Command is the base command
type Command struct {
	// Caller is the optional calling function.
	Caller string `short:"c" long:"caller" description:"The calling function" optional:"yes"`
}
