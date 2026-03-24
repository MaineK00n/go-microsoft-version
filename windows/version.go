package windows

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

// Version represents a Microsoft Windows version
type Version struct {
	Major    int
	Minor    int
	Build    int
	Revision *int
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
		return Version{Major: major, Minor: minor, Build: build}, nil
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
		return Version{Major: major, Minor: minor, Build: build, Revision: &revision}, nil
	default:
		return Version{}, fmt.Errorf("unexpected Windows version format. expected: %q, actual: %q", []string{"<major>.<minor>.<build>", "<major>.<minor>.<build>.<revision>"}, ver)
	}
}

// Compare returns an integer comparing two versions.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
func (v1 Version) Compare(v2 Version) int {
	return cmp.Or(
		cmp.Compare(v1.Major, v2.Major),
		cmp.Compare(v1.Minor, v2.Minor),
		cmp.Compare(v1.Build, v2.Build),
		func() int {
			if v1.Revision == nil || v2.Revision == nil {
				return 0
			}
			return cmp.Compare(*v1.Revision, *v2.Revision)
		}(),
	)
}

// String returns the version string
func (v Version) String() string {
	if v.Revision == nil {
		return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Build)
	}
	return fmt.Sprintf("%d.%d.%d.%d", v.Major, v.Minor, v.Build, *v.Revision)
}
