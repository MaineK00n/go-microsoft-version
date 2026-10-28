package sdk

import (
	"fmt"
	"strings"

	"github.com/MaineK00n/go-microsoft-version/dotnet/core/internal/semver"
)

// Version represents a .NET (Core/5+) SDK version (e.g., "8.0.130", "10.0.402", "11.0.100-rc.3.26522.148").
// Patch holds "<feature band><2-digit patch>", e.g. 402 is feature band 4, patch 2.
type Version struct {
	Major int
	Minor int
	Patch int
	// Prerelease is the dot-separated prerelease identifiers (e.g., "rc.3.26522.148"). Empty for a release version.
	Prerelease string
	// Build is the dot-separated build metadata (e.g., "abc123"). It is ignored in Compare.
	Build string
}

// NewVersion returns a parsed version
func NewVersion(ver string) (Version, error) {
	ver = strings.TrimSpace(ver)
	major, minor, patch, pre, build, err := semver.Parse(ver)
	if err != nil {
		return Version{}, fmt.Errorf("parse .NET (Core/5+) SDK version %q. err: %w", ver, err)
	}
	if patch < 100 {
		return Version{}, fmt.Errorf("unexpected .NET (Core/5+) SDK version format. expected: %q (feature band >= 1), actual: %q", "<major>.<minor>.<feature band><2-digit patch>", ver)
	}
	return Version{Major: major, Minor: minor, Patch: patch, Prerelease: pre, Build: build}, nil
}

// FeatureBand returns the feature band of the SDK version (e.g., 4 for "10.0.402")
func (v Version) FeatureBand() int {
	return v.Patch / 100
}

// BandPatch returns the patch within the feature band (e.g., 2 for "10.0.402")
func (v Version) BandPatch() int {
	return v.Patch % 100
}

// Compare returns an integer comparing two versions.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
// Precedence follows SemVer 2.0.0: a release is newer than its prereleases, and build metadata is ignored.
// Versions in different feature bands are comparable (e.g., 10.0.399 < 10.0.400).
func (v1 Version) Compare(v2 Version) int {
	return semver.Compare(v1.Major, v1.Minor, v1.Patch, v1.Prerelease, v2.Major, v2.Minor, v2.Patch, v2.Prerelease)
}

// String returns the version string
func (v Version) String() string {
	return semver.String(v.Major, v.Minor, v.Patch, v.Prerelease, v.Build)
}
