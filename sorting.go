package gosemver

func less(a, b *Version) bool {
	if a.Major < b.Major {
		return true
	}
	if a.Major == b.Major {
		if a.Minor < b.Minor {
			return true
		}
		if a.Minor == b.Minor {
			if a.Patch < b.Patch {
				return true
			}
			if a.Patch == b.Patch {
				// should it just be alphabetical ordering? well, probably yes
				// semver.org DOES NOT say ANYTHING about sorting pre-release identifiers
				// "Pre-release versions have a lower precedence than the associated normal version," that's all
				return a.Identifiers < b.Identifiers
			}
		}
	}
	return false
}

func (v Versions) Len() int {
	return len(v)
}

func (v Versions) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Versions) Less(i, j int) bool {
	return less(&v[i], &v[j])
}

func (v VersionStrings) Len() int {
	return len(v)
}

func (v VersionStrings) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v VersionStrings) Less(i, j int) bool {
	return less(parseVersion(v[i]), parseVersion(v[j]))
}
