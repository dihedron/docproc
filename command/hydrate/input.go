package hydrate

import "github.com/dihedron/mason/unmarshal"

type Input map[string]interface{}

func (i *Input) UnmarshalFlag(value string) error {
	return unmarshal.FromFlag(value, i)
}
