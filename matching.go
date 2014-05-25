package gosemver

func (v Versions) FindAllOp(operator string, constraint Constraint) Versions {
	result := make(Versions, 0)
	for _, ver := range v {
		if ver.SatisfiesOp(operator, &constraint) {
			result = append(result, ver)
		}
	}
	return result
}

func (v Versions) FindAll(constraint string) (Versions, error) {
	operator, constr, err := parseConstraint(constraint)
	if err != nil {
		return make(Versions, 0), err
	} else {
		return v.FindAllOp(operator, *constr), nil
	}
}
