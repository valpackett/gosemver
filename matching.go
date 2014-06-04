package gosemver

import (
	"fmt"
	"sort"
)

func FindAllOp(vers []Version, operator string, constraint Constraint) []Version {
	var result []Version
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
	}
	return FindAllOp(vers, operator, constr), nil
}

func MustFindAll(vers []Version, constraint string) []Version {
	ver, err := FindAll(vers, constraint)
	if err != nil {
		panic(err)
	}
	return ver
}

func FindMaxOp(vers []Version, operator string, constraint Constraint) (Version, error) {
	matchingVers := FindAllOp(vers, operator, constraint)
	sort.Sort(Versions(matchingVers))
	if len(matchingVers) >= 1 {
		return matchingVers[len(matchingVers)-1], nil
	}
	return Version{}, fmt.Errorf("No matching versions found.")
}

func FindMax(vers []Version, constraint string) (Version, error) {
	operator, constr, err := ParseConstraint(constraint)
	if err != nil {
		return Version{}, err
	}
	return FindMaxOp(vers, operator, constr)
}

func MustFindMax(vers []Version, constraint string) Version {
	ver, err := FindMax(vers, constraint)
	if err != nil {
		panic(err)
	}
	return ver
}
