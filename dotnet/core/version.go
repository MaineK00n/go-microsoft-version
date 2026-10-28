package core

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

// Version represents a .NET (Core/5+) version (e.g., "8.0.2", "9.0.13", "11.0.0-rc.2.26504.105")
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

	rest, build, hasBuild := strings.Cut(ver, "+")
	core, pre, hasPre := strings.Cut(rest, "-")

	ss := strings.Split(core, ".")
	if len(ss) != 3 {
		return Version{}, fmt.Errorf("unexpected .NET (Core/5+) version format. expected: %q, actual: %q", "<major>.<minor>.<patch>[-<prerelease>][+<build>]", ver)
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

	if hasPre {
		if err := validateIdentifiers(pre, true); err != nil {
			return Version{}, fmt.Errorf("parse prerelease. err: %w", err)
		}
	}
	if hasBuild {
		if err := validateIdentifiers(build, false); err != nil {
			return Version{}, fmt.Errorf("parse build metadata. err: %w", err)
		}
	}

	return Version{Major: major, Minor: minor, Patch: patch, Prerelease: pre, Build: build}, nil
}

// validateIdentifiers checks dot-separated identifiers per SemVer 2.0.0.
// Numeric prerelease identifiers must not have leading zeros.
func validateIdentifiers(s string, prerelease bool) error {
	for id := range strings.SplitSeq(s, ".") {
		if id == "" {
			return fmt.Errorf("empty identifier in %q", s)
		}
		for _, c := range id {
			switch {
			case '0' <= c && c <= '9', 'a' <= c && c <= 'z', 'A' <= c && c <= 'Z', c == '-':
			default:
				return fmt.Errorf("invalid character %q in identifier %q", c, id)
			}
		}
		if prerelease && isNumeric(id) && len(id) > 1 && id[0] == '0' {
			return fmt.Errorf("leading zero in numeric identifier %q", id)
		}
	}
	return nil
}

func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// Compare returns an integer comparing two versions.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
// Precedence follows SemVer 2.0.0: a release is newer than its prereleases, and build metadata is ignored.
func (v1 Version) Compare(v2 Version) int {
	return cmp.Or(
		cmp.Compare(v1.Major, v2.Major),
		cmp.Compare(v1.Minor, v2.Minor),
		cmp.Compare(v1.Patch, v2.Patch),
		comparePrerelease(v1.Prerelease, v2.Prerelease),
	)
}

func comparePrerelease(p1, p2 string) int {
	switch {
	case p1 == p2:
		return 0
	case p1 == "":
		return +1
	case p2 == "":
		return -1
	}

	ids1, ids2 := strings.Split(p1, "."), strings.Split(p2, ".")
	for i := range min(len(ids1), len(ids2)) {
		if c := compareIdentifier(ids1[i], ids2[i]); c != 0 {
			return c
		}
	}
	return cmp.Compare(len(ids1), len(ids2))
}

func compareIdentifier(id1, id2 string) int {
	n1, n2 := isNumeric(id1), isNumeric(id2)
	switch {
	case n1 && n2:
		// compare by length first to avoid integer overflow; valid numeric identifiers have no leading zeros
		return cmp.Or(cmp.Compare(len(id1), len(id2)), strings.Compare(id1, id2))
	case n1:
		return -1
	case n2:
		return +1
	default:
		return strings.Compare(id1, id2)
	}
}

// String returns the version string
func (v Version) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.Prerelease != "" {
		sb.WriteString("-" + v.Prerelease)
	}
	if v.Build != "" {
		sb.WriteString("+" + v.Build)
	}
	return sb.String()
}
