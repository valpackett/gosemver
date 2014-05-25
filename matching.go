package gosemver

import (
	"sort"
)

func FindAllOp(vers []Version, operator string, constraint Constraint) []Version {
	result := make([]Version, 0)
	for _, v := range vers {
		if v.SatisfiesOp(operator, &constraint) {
			result = append(result, v)
		}
	}
	return result
}

func FindAll(vers []Version, constraint string) ([]Version, error) {
	operator, constr, err := parseConstraint(constraint)
	if err != nil {
		return make([]Version, 0), err
	} else {
		return FindAllOp(vers, operator, *constr), nil
	}
}

func FindMaxOp(vers []Version, operator string, constraint Constraint) *Version {
	matchingVers := FindAllOp(vers, operator, constraint)
	sort.Sort(Versions(matchingVers))
	return &matchingVers[len(matchingVers)-1]
}

func FindMax(vers []Version, constraint string) (*Version, error) {
	operator, constr, err := parseConstraint(constraint)
	if err != nil {
		return new(Version), err
	} else {
		return FindMaxOp(vers, operator, *constr), nil
	}
}
