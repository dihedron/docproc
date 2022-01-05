package semver

import "github.com/Masterminds/semver/v3"

type SemVer struct {
	Original   string `json:"original,omitempty"`
	Major      uint64 `json:"major"`
	Minor      uint64 `json:"minor"`
	Patch      uint64 `json:"patch"`
	PreRelease string `json:"prerelease,omitempty"`
	Metadata   string `json:"metadata,omitempty"`
}

func NewSemVer(v *semver.Version) *SemVer {
	return &SemVer{
		Original:   v.Original(),
		Major:      v.Major(),
		Minor:      v.Minor(),
		Patch:      v.Patch(),
		PreRelease: v.Prerelease(),
		Metadata:   v.Metadata(),
	}
}
