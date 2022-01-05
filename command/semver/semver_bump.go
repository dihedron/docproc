package semver

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/dihedron/mason/command/base"
)

type Bump struct {
	base.Command
	//Which string `short:"w" long:"which" description:"The field to bump." choice:"major" choice:"M" choice:"minor" choice:"m" choice:"patch" choice:"p" choice:"revision" choice:"rev" choice:"r" optional:"yes" default:"patch"`
	Major    bool `short:"M" long:"major" description:"Bump the major field." optional:"yes"`
	Minor    bool `short:"m" long:"minor" description:"Bump the minor field." optional:"yes"`
	Patch    bool `short:"p" long:"patch" description:"Bump the patch (aka revision) field." optional:"yes"`
	Revision bool `short:"r" long:"revision" description:"Bump the revision (aka patch) field." optional:"yes"`
}

func (cmd *Bump) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("at least one version must be provided")
	}

	versions := []*SemVer{}

	for _, arg := range args {
		v, err := semver.NewVersion(arg) //"1.2.3-beta.1+build345")
		if err != nil {
			return fmt.Errorf("error parsing version '%s': %w", arg, err)
		}
		cmd.Patch = cmd.Patch || cmd.Revision
		if (cmd.Major && cmd.Minor) || (cmd.Major && cmd.Patch) || (cmd.Minor && cmd.Patch) {
			return fmt.Errorf("only one of --major (value: %t), --minor (value: %t) and --patch (value: %t) must be provided at once", cmd.Major, cmd.Minor, cmd.Patch)
		}
		if !cmd.Major && !cmd.Minor && !cmd.Patch {
			// no flag set, default to...
			cmd.Patch = true
		}
		if cmd.Major {
			v1 := v.IncMajor()
			v = &v1
		} else if cmd.Minor {
			v1 := v.IncMinor()
			v = &v1
		} else if cmd.Patch {
			v1 := v.IncPatch()
			v = &v1
		}
		if cmd.Automation {
			versions = append(versions, NewSemVer(v))
		} else {
			fmt.Printf("%s\n", v.String())
		}
	}

	if cmd.Automation {
		var (
			data []byte
			err  error
		)
		if len(versions) > 1 {
			data, err = json.Marshal(versions)
		} else {
			data, err = json.Marshal(versions[0])
		}
		if err != nil {
			return fmt.Errorf("error marshalling version(s) '%v' to JSON: %w", versions, err)
		}
		fmt.Printf("%s", string(data))
	}

	return nil
}
