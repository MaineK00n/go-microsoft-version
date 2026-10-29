// Package semver provides SemVer 2.0.0 parsing and comparison shared by the .NET (Core/5+) runtime and SDK versions.
package semver

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

// Parse parses "<major>.<minor>.<patch>[-<prerelease>][+<build>]".
// Leading zeros in major, minor and patch are accepted (e.g. "08.0.02" is 8.0.2), as in the
// earlier dotnet/core package. Errors do not repeat ver; callers wrap them with it.
func Parse(ver string) (major, minor, patch int, prerelease, build string, err error) {
	rest, build, hasBuild := strings.Cut(ver, "+")
	core, prerelease, hasPre := strings.Cut(rest, "-")

	ss := strings.Split(core, ".")
	if len(ss) != 3 {
		return 0, 0, 0, "", "", fmt.Errorf("unexpected version format. expected: %q", "<major>.<minor>.<patch>[-<prerelease>][+<build>]")
	}

	major, err = strconv.Atoi(ss[0])
	if err != nil {
		return 0, 0, 0, "", "", fmt.Errorf("parse major version. err: %w", err)
	}
	minor, err = strconv.Atoi(ss[1])
	if err != nil {
		return 0, 0, 0, "", "", fmt.Errorf("parse minor version. err: %w", err)
	}
	patch, err = strconv.Atoi(ss[2])
	if err != nil {
		return 0, 0, 0, "", "", fmt.Errorf("parse patch version. err: %w", err)
	}

	if hasPre {
		if err := validateIdentifiers(prerelease, true); err != nil {
			return 0, 0, 0, "", "", fmt.Errorf("parse prerelease. err: %w", err)
		}
	}
	if hasBuild {
		if err := validateIdentifiers(build, false); err != nil {
			return 0, 0, 0, "", "", fmt.Errorf("parse build metadata. err: %w", err)
		}
	}

	return major, minor, patch, prerelease, build, nil
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

// Compare returns an integer comparing two versions by SemVer 2.0.0 precedence.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
// Build metadata is not taken as an argument since it is ignored for precedence.
func Compare(major1, minor1, patch1 int, prerelease1 string, major2, minor2, patch2 int, prerelease2 string) int {
	return cmp.Or(
		cmp.Compare(major1, major2),
		cmp.Compare(minor1, minor2),
		cmp.Compare(patch1, patch2),
		ComparePrerelease(prerelease1, prerelease2),
	)
}

// ComparePrerelease compares two dot-separated prerelease strings.
// An empty string means a release, which ranks above any prerelease.
func ComparePrerelease(p1, p2 string) int {
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

// String returns "<major>.<minor>.<patch>[-<prerelease>][+<build>]".
func String(major, minor, patch int, prerelease, build string) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%d.%d.%d", major, minor, patch)
	if prerelease != "" {
		sb.WriteString("-" + prerelease)
	}
	if build != "" {
		sb.WriteString("+" + build)
	}
	return sb.String()
}
