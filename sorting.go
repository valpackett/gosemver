package gosemver

func (v Versions) Len() int {
	return len(v)
}

func (v Versions) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Versions) Less(i, j int) bool {
	// breaks (?) if you have, like, version 8.131072.65536. who has versions like these though?
	iNum := v[i].Major*1000000 + v[i].Minor*1000 + v[i].Patch
	jNum := v[j].Major*1000000 + v[j].Minor*1000 + v[j].Patch
	if iNum == jNum {
		// should it just be alphabetical ordering? well, probably yes
		// semver.org DOES NOT say ANYTHING about sorting pre-release identifiers
		// "Pre-release versions have a lower precedence than the associated normal version," that's all
		return v[i].Identifiers < v[j].Identifiers
	}
	return iNum < jNum
}
