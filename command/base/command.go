package base

import (
	"os"
	"regexp"
	"strings"

	"go.uber.org/zap"
)

// Command collects the base set of options in common to all commands.
type Command struct {
	// Automation enables automation-friendly JSON output.
	Automation bool `short:"A" long:"automation" description:"Whether to output in automation friendly JSON format." optional:"yes"`
	// Parameters are a set of <key>:<value> or <key>=<value> pairs, that ca be used for substitution in inputs.
	Parameters []Parameter `short:"P" long:"parameter" description:"A set of parameters, in <key>:<value> format." optional:"yes"`
}

// BindVariables provides a way to bind variables in strings with
// parameters provided either in the environment, or on the command
// line as '--parameter' pairs. In order to bind all variables in a
// string do as follows:
//   cmd.Bind("{env:SOME_ENV_VAR}={cli:key1}+{cli:key2}")
// and it will be bound according to the following rules:
// * variables of the form {env:<ENVVAR>} are bound to the corresponding
//   environment variable
// * variables of the forma {cli:<PARAM>} are bound to the corresponding
//   parameter as provided on the command line via the command line switch.
func (cmd *Command) BindVariables(value string) string {
	result := value
	re := regexp.MustCompile(`(?:\{(env|cli)\:([a-zA-Z0-9-_@\.]+)\})`)
	groups := re.FindAllStringSubmatch(value, -1)
	if groups == nil {
		// no match, no need to bind
		return ""
	}
	for _, group := range groups {
		zap.S().Debugf("variable match: '%s' '%s' '%s'\n", group[0], group[1], group[2])
		switch group[1] {
		case "cli":
			for _, parameter := range cmd.Parameters {
				if group[2] == parameter.Key {
					zap.S().Debugf("replacing variable '%s' with value '%s' in '%s'\n", group[0], parameter.Value, result)
					result = strings.ReplaceAll(result, group[0], parameter.Value)
					continue
				}
			}
		case "env":
			e := os.Getenv(group[2])
			zap.S().Debugf("replacing variable '%s' with value '%s' in '%s'\n", group[0], e, result)
			result = strings.ReplaceAll(result, group[0], e)
		}
	}
	zap.S().Debugf("'%s'\n", result)
	return result
}
