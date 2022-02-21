package unmarshal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

type DataFormat uint8

const (
	DataFormatJSON DataFormat = iota
	DataFormatYAML
)

// FromFlag unmarshals a command line flag into an argument;
// if the command line argument starts with a '@' it is assumed to
// be a file on the local filesystem, it is read into memory and then
// unmarshalled into the object struct, which must be appropriately
// annotated; if it does not start with '@', it is assumed to be an
// inline JSON representation and is unmarshalled as such.
func FromFlag(value string) (interface{}, error) {
	format := DataFormatJSON
	var content []byte
	if strings.HasPrefix(value, "@") {
		filename := strings.TrimPrefix(value, "@")
		info, err := os.Stat(filename)
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file '%s' does not exist: %w", filename, err)
		}
		if info.IsDir() {
			return nil, fmt.Errorf("'%s' is a directory, not a file", filename)
		}
		content, err = ioutil.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("error reading file '%s': %w", filename, err)
		}
		ext := path.Ext(filename)
		switch ext {
		case ".yaml", ".yml":
			format = DataFormatYAML
		case ".json":
			format = DataFormatJSON
		default:
			return nil, fmt.Errorf("unsupported data format in file: %s", path.Ext(filename))
		}
	} else {
		value = strings.TrimSpace(value)
		content = []byte(value)
		if strings.HasPrefix(value, "---") {
			format = DataFormatYAML
		} else if strings.HasPrefix(value, "{") {
			format = DataFormatJSON
		} else {
			return nil, fmt.Errorf("unrecognisable input format on STDIN")
		}
	}
	switch format {
	case DataFormatJSON:
		// first attempt unmarshalling to a map (like a struct would)...
		m := map[string]interface{}{}
		if err := json.Unmarshal(content, &m); err != nil {
			if err, ok := err.(*json.UnmarshalTypeError); ok {
				if err.Value == "array" && err.Offset == 1 {
					// it is not a struct, it's an array, let's try that...
					a := []interface{}{}
					if err := yaml.Unmarshal(content, &a); err != nil {
						return nil, fmt.Errorf("error unmarshalling from JSON: %w", err)
					}
					return a, nil
				}
			}
			return nil, fmt.Errorf("error unmarshalling from JSON: %w", err)
		}
		return m, nil
	case DataFormatYAML:
		object := map[string]interface{}{}
		if err := yaml.Unmarshal(content, object); err != nil {
			return nil, fmt.Errorf("error unmarshalling from YAML: %w", err)
		}
		return object, nil
	default:
		return nil, fmt.Errorf("unsupported encoding: %v", format)
	}
}

func toJSON(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "  ")
	return string(data)
}
