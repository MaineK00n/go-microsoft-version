package visualstudio

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

// VersionType represents the version format type of Visual Studio
type VersionType int

const (
	// Legacy represents Visual Studio 2015 and earlier (4-segment: e.g., "14.0.27552.0", "12.0.40700.0")
	Legacy VersionType = iota
	// Modern represents Visual Studio 2017 and later (3-segment: e.g., "17.8.6", "15.9.38")
	Modern
)

// Version represents a Microsoft Visual Studio version
type Version struct {
	Type     VersionType
	Major    int
	Minor    int
	Build    int
	Revision int
}

// NewVersion returns a parsed version
func NewVersion(ver string) (Version, error) {
	ver = strings.TrimSpace(ver)
	switch ss := strings.Split(ver, "."); len(ss) {
	case 3:
		major, err := strconv.Atoi(ss[0])
		if err != nil {
			return Version{}, fmt.Errorf("parse major version. err: %w", err)
		}
		minor, err := strconv.Atoi(ss[1])
		if err != nil {
			return Version{}, fmt.Errorf("parse minor version. err: %w", err)
		}
		build, err := strconv.Atoi(ss[2])
		if err != nil {
			return Version{}, fmt.Errorf("parse build version. err: %w", err)
		}
		return Version{Type: Modern, Major: major, Minor: minor, Build: build}, nil
	case 4:
		major, err := strconv.Atoi(ss[0])
		if err != nil {
			return Version{}, fmt.Errorf("parse major version. err: %w", err)
		}
		minor, err := strconv.Atoi(ss[1])
		if err != nil {
			return Version{}, fmt.Errorf("parse minor version. err: %w", err)
		}
		build, err := strconv.Atoi(ss[2])
		if err != nil {
			return Version{}, fmt.Errorf("parse build version. err: %w", err)
		}
		revision, err := strconv.Atoi(ss[3])
		if err != nil {
			return Version{}, fmt.Errorf("parse revision version. err: %w", err)
		}
		return Version{Type: Legacy, Major: major, Minor: minor, Build: build, Revision: revision}, nil
	default:
		return Version{}, fmt.Errorf("unexpected Visual Studio version format. expected: %q, actual: %q", []string{"<major>.<minor>.<build>", "<major>.<minor>.<build>.<revision>"}, ver)
	}
}

// Compare returns an integer comparing two versions.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
// Modern versions are always newer than Legacy versions.
func (v1 Version) Compare(v2 Version) int {
	if v1.Type != v2.Type {
		return cmp.Compare(v1.Type, v2.Type)
	}
	return cmp.Or(
		cmp.Compare(v1.Major, v2.Major),
		cmp.Compare(v1.Minor, v2.Minor),
		cmp.Compare(v1.Build, v2.Build),
		cmp.Compare(v1.Revision, v2.Revision),
	)
}

// String returns the version string
func (v Version) String() string {
	switch v.Type {
	case Legacy:
		return fmt.Sprintf("%d.%d.%d.%d", v.Major, v.Minor, v.Build, v.Revision)
	default:
		return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Build)
	}
}
