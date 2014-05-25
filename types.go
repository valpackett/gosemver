package gosemver

type Version struct {
	Prefix        string // Technically not a part of semver
	Major         int
	Minor         int
	Patch         int
	Identifiers   string
	BuildMetadata string
}

type Versions []Version

type Constraint struct {
	Major      int
	Minor      int
	Patch      int
	MatchPatch bool // because Go does not have neither uninitialized variables nor Option/Maybe, how the fuck else do we know if Patch == 0 means match against zero or don't match?
}
