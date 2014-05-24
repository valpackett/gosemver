package gosemver

import (
	"fmt"
)

// implementing fmt.Stringer
func (v *Version) String() string {
	out := fmt.Sprintf("%s%d.%d.%d", v.Prefix, v.Major, v.Minor, v.Patch)
	if v.Identifiers != "" {
		out += "-" + v.Identifiers
	}
	if v.BuildMetadata != "" {
		out += "+" + v.BuildMetadata
	}
	return out
}
