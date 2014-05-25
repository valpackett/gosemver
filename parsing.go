package gosemver

import (
	"regexp"
	"strconv"
)

var constrRegexp = regexp.MustCompile(`^(|~>|\^|<|>|<=|>=|==) ?([0-9]+)\.([0-9]+)(\.([0-9]+))?`)

func parseConstraint(input string) (string, *Constraint) {
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
	}
	return operator, constr
}
