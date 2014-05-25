package gosemver

import (
	"regexp"
	"strconv"
)

func (v *Version) SatisfiesPessimistic(c *Constraint) bool {
	if c.MatchPatch {
		return v.Major == c.Major && v.Minor == c.Minor && v.Patch >= c.Patch
	}
	return v.Major == c.Major && v.Minor >= c.Minor
}

func (v *Version) SatisfiesExact(c *Constraint) bool {
	result := v.Major == c.Major && v.Minor == c.Minor
	if c.MatchPatch {
		return result && v.Patch == c.Patch
	}
	return result
}

func (v *Version) SatisfiesLessThan(c *Constraint) bool {
	if v.Major < c.Major {
		return true
	}
	if v.Minor < c.Minor {
		return true
	}
	if c.MatchPatch && v.Patch < c.Patch {
		return true
	}
	return false
}

func (v *Version) SatisfiesLessThanOrEqual(c *Constraint) bool {
	return v.SatisfiesLessThan(c) || v.SatisfiesExact(c)
}

func (v *Version) SatisfiesGreaterThan(c *Constraint) bool {
	if v.Major > c.Major {
		return true
	}
	if v.Minor > c.Minor {
		return true
	}
	if c.MatchPatch && v.Patch > c.Patch {
		return true
	}
	return false
}

func (v *Version) SatisfiesGreaterThanOrEqual(c *Constraint) bool {
	return v.SatisfiesGreaterThan(c) || v.SatisfiesExact(c)
}

var constrRegexp = regexp.MustCompile(`^(|~>|\^|<|>|<=|>=|==) ?([0-9]+)\.([0-9]+)(\.([0-9]+))?`)

func (v *Version) Satisfies(constraint string) bool {
	if constraint == "" || constraint == "*" || constraint == "x" {
		return true
	}

	operator := ""
	constr := new(Constraint)
	matches := constrRegexp.FindStringSubmatch(constraint)
	if matches != nil {
		operator = matches[1]
		constr.Major, _ = strconv.Atoi(matches[2])
		constr.Minor, _ = strconv.Atoi(matches[3])
		if matches[5] != "" {
			constr.Patch, _ = strconv.Atoi(matches[5])
			constr.MatchPatch = true
		}
	}

	// fmt.Println(operator, constr)

	if operator == "~>" || operator == "^" {
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
