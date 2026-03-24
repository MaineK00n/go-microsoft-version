package edge

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

// VersionType represents the rendering engine type of Microsoft Edge
type VersionType int

const (
	// EdgeHTML represents the legacy EdgeHTML-based versions (e.g., "20.10240", "44.17763")
	EdgeHTML VersionType = iota
	// Chromium represents the Chromium-based versions (e.g., "88.0.705.18", "146.0.3856.13")
	Chromium
)

// Version represents a Microsoft Edge version
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
	case 2:
		major, err := strconv.Atoi(ss[0])
		if err != nil {
			return Version{}, fmt.Errorf("parse major version. err: %w", err)
		}

		build, err := strconv.Atoi(ss[1])
		if err != nil {
			return Version{}, fmt.Errorf("parse build version. err: %w", err)
		}

		return Version{Type: EdgeHTML, Major: major, Build: build}, nil
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

		return Version{Type: Chromium, Major: major, Minor: minor, Build: build, Revision: revision}, nil
	default:
		return Version{}, fmt.Errorf("unexpected Edge version format. expected: %q, actual: %q", []string{"<major>.<build>", "<major>.<minor>.<build>.<revision>"}, ver)
	}
}

// Compare returns an integer comparing two versions.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
// Chromium-based versions are always newer than EdgeHTML-based versions.
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

// String returns the full version string
func (v Version) String() string {
	switch v.Type {
	case EdgeHTML:
		return fmt.Sprintf("%d.%d", v.Major, v.Build)
	default:
		return fmt.Sprintf("%d.%d.%d.%d", v.Major, v.Minor, v.Build, v.Revision)
	}
}
