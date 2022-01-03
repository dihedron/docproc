package command

import (
	"strings"
)

type Template struct {
	Path string
	Main bool
}

// --template=@tests/outer.tpl for main template
func (t *Template) UnmarshalFlag(value string) error {
	value = strings.TrimLeft(value, "\r\n\t ")
	if strings.HasPrefix(value, "@") {
		t.Path = strings.TrimPrefix(value, "@")
		t.Main = true
	} else {
		t.Path = value
	}
	return nil
}
