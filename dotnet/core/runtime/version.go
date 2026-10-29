package runtime

import (
	"fmt"
	"strings"

	"github.com/MaineK00n/go-microsoft-version/dotnet/core/internal/semver"
)

// Version represents a .NET (Core/5+) runtime version (e.g., "8.0.2", "9.0.13", "11.0.0-rc.2.26504.105")
type Version struct {
	Major int
	Minor int
	Patch int
	// Prerelease is the dot-separated prerelease identifiers (e.g., "rc.2.26504.105"). Empty for a release version.
	Prerelease string
	// Build is the dot-separated build metadata (e.g., "abc123"). It is ignored in Compare.
	Build string
}

// NewVersion returns a parsed version
func NewVersion(ver string) (Version, error) {
	ver = strings.TrimSpace(ver)
	major, minor, patch, pre, build, err := semver.Parse(ver)
	if err != nil {
		return Version{}, fmt.Errorf("parse .NET (Core/5+) runtime version %q. err: %w", ver, err)
	}
	return Version{Major: major, Minor: minor, Patch: patch, Prerelease: pre, Build: build}, nil
}

// Compare returns an integer comparing two versions.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
// Precedence follows SemVer 2.0.0: a release is newer than its prereleases, and build metadata is ignored.
func (v1 Version) Compare(v2 Version) int {
	return semver.Compare(v1.Major, v1.Minor, v1.Patch, v1.Prerelease, v2.Major, v2.Minor, v2.Patch, v2.Prerelease)
}

// String returns the version string
func (v Version) String() string {
	return semver.String(v.Major, v.Minor, v.Patch, v.Prerelease, v.Build)
}
