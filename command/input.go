package command

import "github.com/dihedron/ginkgo/unmarshal"

type Input map[string]interface{}

func (i *Input) UnmarshalFlag(value string) error {
	return unmarshal.FromFlag(value, i)
}
