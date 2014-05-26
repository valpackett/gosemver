package gosemver

import (
	"fmt"
	"sort"
)

func FindAllOp(vers []Version, operator string, constraint Constraint) []Version {
	result := make([]Version, 0)
	for _, v := range vers {
		if v.SatisfiesOp(operator, constraint) {
			result = append(result, v)
		}
	}
	return result
}

func FindAll(vers []Version, constraint string) ([]Version, error) {
	operator, constr, err := ParseConstraint(constraint)
	if err != nil {
		return make([]Version, 0), err
	} else {
		return FindAllOp(vers, operator, constr), nil
	}
}

func FindMaxOp(vers []Version, operator string, constraint Constraint) (Version, error) {
	matchingVers := FindAllOp(vers, operator, constraint)
	sort.Sort(Versions(matchingVers))
	if len(matchingVers) >= 1 {
		return matchingVers[len(matchingVers)-1], nil
	} else {
		return Version{}, fmt.Errorf("No matching versions found.")
	}
}

func FindMax(vers []Version, constraint string) (Version, error) {
	operator, constr, err := ParseConstraint(constraint)
	if err != nil {
		return Version{}, err
	} else {
		return FindMaxOp(vers, operator, constr)
	}
}
