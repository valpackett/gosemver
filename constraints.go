package gosemver

func (v Version) SatisfiesPessimistic(c Constraint) bool {
	if c.MatchPatch {
		return v.Major == c.Major && v.Minor == c.Minor && v.Patch >= c.Patch
	}
	return v.Major == c.Major && v.Minor >= c.Minor
}

func (v Version) SatisfiesExact(c Constraint) bool {
	result := v.Major == c.Major && v.Minor == c.Minor
	if c.MatchPatch {
		return result && v.Patch == c.Patch
	}
	return result
}

func (v Version) SatisfiesLessThan(c Constraint) bool {
	if v.Major < c.Major {
		return true
	}
	if v.Major == c.Major {
		if v.Minor < c.Minor {
			return true
		}
		if v.Minor == c.Minor && c.MatchPatch && v.Patch < c.Patch {
			return true
		}
	}
	return false
}

func (v Version) SatisfiesLessThanOrEqual(c Constraint) bool {
	return v.SatisfiesLessThan(c) || v.SatisfiesExact(c)
}

func (v Version) SatisfiesGreaterThan(c Constraint) bool {
	if v.Major > c.Major {
		return true
	}
	if v.Major == c.Major {
		if v.Minor > c.Minor {
			return true
		}
		if v.Minor == c.Minor && c.MatchPatch && v.Patch > c.Patch {
			return true
		}
	}
	return false
}

func (v Version) SatisfiesGreaterThanOrEqual(c Constraint) bool {
	return v.SatisfiesGreaterThan(c) || v.SatisfiesExact(c)
}

func (v Version) Satisfies(constraint string) (result bool, err error) {
	if constraint == "" || constraint == "*" || constraint == "x" {
		return true, nil
	}
	operator, constr, err := parseConstraint(constraint)
	if err != nil {
		return false, err
	} else {
		return v.SatisfiesOp(operator, constr), nil
	}
}

func (v Version) SatisfiesOp(operator string, constr Constraint) bool {
	if operator == "*" || operator == "x" {
		return true
	} else if operator == "~>" || operator == "^" {
		return v.SatisfiesPessimistic(constr)
	} else if operator == "" || operator == "==" {
		return v.SatisfiesExact(constr)
	} else if operator == "<" {
		return v.SatisfiesLessThan(constr)
	} else if operator == ">" {
		return v.SatisfiesGreaterThan(constr)
	} else if operator == "<=" {
		return v.SatisfiesLessThanOrEqual(constr)
	} else if operator == ">=" {
		return v.SatisfiesGreaterThanOrEqual(constr)
	}
	return false
}
