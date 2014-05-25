package gosemver

func (v Versions) Len() int {
	return len(v)
}

func (v Versions) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Versions) Less(i, j int) bool {
	if v[i].Major < v[j].Major {
		return true
	}
	if v[i].Major == v[j].Major {
		if v[i].Minor < v[j].Minor {
			return true
		}
		if v[i].Minor == v[j].Minor {
			if v[i].Patch < v[j].Patch {
				return true
			}
			if v[i].Patch == v[j].Patch {
				// should it just be alphabetical ordering? well, probably yes
				// semver.org DOES NOT say ANYTHING about sorting pre-release identifiers
				// "Pre-release versions have a lower precedence than the associated normal version," that's all
				return v[i].Identifiers < v[j].Identifiers
			}
		}
	}
	return false
}
