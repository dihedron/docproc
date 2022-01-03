package base

import (
	"fmt"
	"regexp"
)

// Parameter represents a parameter that can be used to
// values in configuration files.
type Parameter struct {
	Key   string
	Value string
}

func (p *Parameter) UnmarshalFlag(value string) error {
	re := regexp.MustCompile(`^([a-zA-Z0-9-_@\.]+)(?:\:|=)(.*)$`)
	matches := re.FindStringSubmatch(value)
	if matches == nil {
		return fmt.Errorf("invalid format for parameter '%s'", value)
	}
	p.Key = matches[1]
	p.Value = matches[2]
	return nil
}
