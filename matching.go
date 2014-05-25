package gosemver

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
