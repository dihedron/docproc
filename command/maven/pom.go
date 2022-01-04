package maven

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Project struct {
	XMLName        xml.Name `xml:"project" json:"project,omitempty"`
	XMLNS          string   `xml:"xmlns,attr" json:"xmlns,omitempty"`
	XSI            string   `xml:"xsi,attr" json:"xsi,omitempty"`
	SchemaLocation string   `xml:"schemaLocation,attr" json:"schemaLocation,omitempty"`
	ModelVersion   struct {
		Text string `xml:",chardata" json:"text,omitempty"`
	} `xml:"modelVersion" json:"modelversion,omitempty"`
	Parent struct {
		// Text    string `xml:",chardata" json:"text,omitempty"`
		GroupId struct {
			Text string `xml:",chardata" json:"text,omitempty"`
		} `xml:"groupId" json:"groupid,omitempty"`
		ArtifactId struct {
			Text string `xml:",chardata" json:"text,omitempty"`
		} `xml:"artifactId" json:"artifactid,omitempty"`
		Version struct {
			Text string `xml:",chardata" json:"text,omitempty"`
		} `xml:"version" json:"version,omitempty"`
		Packaging struct {
			Text string `xml:",chardata" json:"text,omitempty"`
		} `xml:"packaging" json:"packaging,omitempty"`
		RelativePath struct {
			Text string `xml:",chardata" json:"text,omitempty"`
		} `xml:"relativePath" json:"relativePath,omitempty"`
	} `xml:"parent" json:"parent,omitempty"`
	GroupId struct {
		Text string `xml:",chardata" json:"text,omitempty"`
	} `xml:"groupId" json:"groupid,omitempty"`
	ArtifactId struct {
		Text string `xml:",chardata" json:"text,omitempty"`
	} `xml:"artifactId" json:"artifactid,omitempty"`
	Version struct {
		Text string `xml:",chardata" json:"text,omitempty"`
	} `xml:"version" json:"version,omitempty"`
	Packaging struct {
		Text string `xml:",chardata" json:"text,omitempty"`
	} `xml:"packaging" json:"packaging,omitempty"`
	Modules struct {
		// Text   string `xml:",chardata" json:"text,omitempty"`
		Module []struct {
			Text string `xml:",chardata" json:"text,omitempty"`
		} `xml:"module" json:"module,omitempty"`
	} `xml:"modules" json:"modules,omitempty"`
}

func (p *Project) UnmarshalFlag(value string) error {
	var content []byte
	if strings.HasPrefix(value, "@") {
		filename := strings.TrimPrefix(value, "@")
		info, err := os.Stat(filename)
		if os.IsNotExist(err) {
			return fmt.Errorf("file '%s' does not exist: %w", filename, err)
		}
		if info.IsDir() {
			return fmt.Errorf("'%s' is a directory, not a file", filename)
		}
		content, err = ioutil.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("error reading file '%s': %w", filename, err)
		}
	} else {
		content = []byte(value)
	}

	return xml.Unmarshal(content, &p)
}
