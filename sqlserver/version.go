package sqlserver

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

// Version represents a Microsoft SQL Server version (e.g., "15.0.2095.3", "16.0.4165.4")
type Version struct {
	Major    int
	Minor    int
	Build    int
	Revision int
}

// NewVersion returns a parsed version
func NewVersion(ver string) (Version, error) {
	ver = strings.TrimSpace(ver)
	ss := strings.Split(ver, ".")
	if len(ss) != 4 {
		return Version{}, fmt.Errorf("unexpected Microsoft SQL Server version format. expected: %q, actual: %q", "<major>.<minor>.<build>.<revision>", ver)
	}

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

	return Version{Major: major, Minor: minor, Build: build, Revision: revision}, nil
}

// Compare returns an integer comparing two versions.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
func (v1 Version) Compare(v2 Version) int {
	return cmp.Or(
		cmp.Compare(v1.Major, v2.Major),
		cmp.Compare(v1.Minor, v2.Minor),
		cmp.Compare(v1.Build, v2.Build),
		cmp.Compare(v1.Revision, v2.Revision),
	)
}

// String returns the version string
func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", v.Major, v.Minor, v.Build, v.Revision)
}
