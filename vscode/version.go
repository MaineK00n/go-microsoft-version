package vscode

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

// Version represents a Visual Studio Code version (e.g., "1.100.1", "1.104.0")
type Version struct {
	Major int
	Minor int
	Patch int
}

// NewVersion returns a parsed version
func NewVersion(ver string) (Version, error) {
	ver = strings.TrimSpace(ver)
	ss := strings.Split(ver, ".")
	if len(ss) != 3 {
		return Version{}, fmt.Errorf("unexpected VS Code version format. expected: %q, actual: %q", "<major>.<minor>.<patch>", ver)
	}

	major, err := strconv.Atoi(ss[0])
	if err != nil {
		return Version{}, fmt.Errorf("parse major version. err: %w", err)
	}
	minor, err := strconv.Atoi(ss[1])
	if err != nil {
		return Version{}, fmt.Errorf("parse minor version. err: %w", err)
	}
	patch, err := strconv.Atoi(ss[2])
	if err != nil {
		return Version{}, fmt.Errorf("parse patch version. err: %w", err)
	}

	return Version{Major: major, Minor: minor, Patch: patch}, nil
}

// Compare returns an integer comparing two versions.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
func (v1 Version) Compare(v2 Version) int {
	return cmp.Or(
		cmp.Compare(v1.Major, v2.Major),
		cmp.Compare(v1.Minor, v2.Minor),
		cmp.Compare(v1.Patch, v2.Patch),
	)
}

// String returns the version string
func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
