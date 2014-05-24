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
