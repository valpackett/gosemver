package gosemver

import (
	"regexp"
	"strconv"
)

func (v *Version) SatisfiesPessimistic(major, minor int) bool {
	return v.Major == major && v.Minor >= minor
}

func (v *Version) SatisfiesPessimisticWithPatch(major, minor, patch int) bool {
	return v.Major == major && v.Minor == minor && v.Patch >= patch
}

func (v *Version) Satisfies(constraint string) bool {
	pessimisticWithPatchMatches := regexp.MustCompile(`^(?:~>|\^) ?([0-9]+)\.([0-9]+)(\.([0-9]+))`).FindStringSubmatch(constraint)
	if pessimisticWithPatchMatches != nil {
		major, _ := strconv.Atoi(pessimisticWithPatchMatches[1])
		minor, _ := strconv.Atoi(pessimisticWithPatchMatches[2])
		patch, _ := strconv.Atoi(pessimisticWithPatchMatches[4])
		return v.SatisfiesPessimisticWithPatch(major, minor, patch)
	}
	pessimisticMatches := regexp.MustCompile(`^(?:~>|\^) ?([0-9]+)\.([0-9]+)`).FindStringSubmatch(constraint)
	if pessimisticMatches != nil {
		major, _ := strconv.Atoi(pessimisticMatches[1])
		minor, _ := strconv.Atoi(pessimisticMatches[2])
		return v.SatisfiesPessimistic(major, minor)
	}

	return false
}
