package gosemver

import (
	"fmt"
	"regexp"
	"strconv"
)

var constrRegexp = regexp.MustCompile(`^(|~>|\^|<|>|<=|>=|==) ?([0-9]+)\.([0-9]+)(\.([0-9]+))?`)

func parseConstraint(input string) (string, *Constraint, error) {
	operator := ""
	constr := new(Constraint)
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

var verRegexp = regexp.MustCompile(`^([^0-9]*)([0-9]+)\.([0-9]+)\.([0-9]+)(\-([^+]+))?(\+(.*))?`)

func parseVersion(input string) (*Version, error) {
	ver := new(Version)
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

func parseVersions(input []string) ([]Version, error) {
	vers := make([]Version, len(input))
	for i, verStr := range input {
		parsedVersion, err := parseVersion(verStr)
		if err != nil {
			return vers, err
		}
		vers[i] = *parsedVersion
	}
	return vers, nil
}
