package gosemver

import (
	"fmt"
	"regexp"
	"strconv"
)

var constrRegexp = regexp.MustCompile(`^(|~>|\^|<|>|<=|>=|==) ?([0-9]+)\.([0-9]+)(\.([0-9]+))?`)

func ParseConstraint(input string) (string, Constraint, error) {
	operator := ""
	constr := Constraint{}
	matches := constrRegexp.FindStringSubmatch(input)
	if matches != nil {
		operator = matches[1]
		constr.Major, _ = strconv.Atoi(matches[2])
		constr.Minor, _ = strconv.Atoi(matches[3])
		if matches[5] != "" {
			constr.Patch, _ = strconv.Atoi(matches[5])
			constr.MatchPatch = true
		}
	} else {
		return operator, constr, fmt.Errorf("Invalid constraint string: %s", input)
	}
	return operator, constr, nil
}

func MustParseConstraint(input string) (string, constraint) {
        operator, constr, err := ParseConstraint(input)
        if err != nil {
                panic(err)
        }
        return operator, constr
}

var verRegexp = regexp.MustCompile(`^([^0-9]*)([0-9]+)\.([0-9]+)\.([0-9]+)(\-([^+]+))?(\+(.*))?`)

func ParseVersion(input string) (Version, error) {
	ver := Version{}
	matches := verRegexp.FindStringSubmatch(input)
	if matches != nil {
		ver.Prefix = matches[1]
		ver.Major, _ = strconv.Atoi(matches[2])
		ver.Minor, _ = strconv.Atoi(matches[3])
		ver.Patch, _ = strconv.Atoi(matches[4])
		ver.Identifiers = matches[6]
		ver.BuildMetadata = matches[8]
	} else {
		return ver, fmt.Errorf("Invalid version string: %s", input)
	}
	return ver, nil
}

func MustParseVersion(input string) Version {
        ver, err := ParseVersion(input)
        if err != nil {
                panic(err)
        }
        return ver
}

func ParseVersions(input []string) ([]Version, error) {
	vers := make([]Version, len(input))
	for i, verStr := range input {
		var err error
		vers[i], err = ParseVersion(verStr)
		if err != nil {
			return vers, err
		}
	}
	return vers, nil
}

func MustParseVersions(input []string) []Version {
        vers, err := ParseVersions(input)
        if err != nil {
                panic(err)
        }
        return vers
}
