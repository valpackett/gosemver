# gosemver

A [Semantic Versioning](http://semver.org) library for the Go programming language.

## Usage

```go
import "github.com/myfreeweb/gosemver"
```

Parsing:

```go
gosemver.parseVersion("v0.1.0-alpha+build001") // gosemver.Version{"v", 0, 1, 0, "alpha", "build001"}, nil
gosemver.parseVersion("AAAAAAA") // gosemver.Version{"", 0, 0, 0, "", ""}, error

gosemver.parseVersions([]string{"0.1.0", "0.2.0"}) // []gosemver.Version{{"", 0, 1, 0, "", ""}, {"", 0, 2, 0, "", ""},}, nil
```

Sorting:

```go
import "sort"

vers := []gosemver.Version{
  gosemver.Version{"v", 1, 0, 0, "", ""},
  gosemver.Version{"v", 0, 1, 0, "", ""},
}

sort.Sort(gosemver.Versions(vers))
```

Output:

```go
ver := gosemver.Version{"v", 2, 3, 1, "alpha", "build.001"}
ver.String() // "v2.3.1-alpha+build.001"
```

Constraints:

```go
ver := gosemver.Version{"", 3, 0, 3, "", ""}
// returns result, error:
ver.Satisfies("*") // true, nil
ver.Satisfies("== 3.0.3") // true, nil
ver.Satisfies(">= 3.0.1") // true, nil
ver.Satisfies(">= 3.0") // true, nil
ver.Satisfies("> 3.0.0") // true, nil
ver.Satisfies("~> 3.0.4") // false, nil
ver.Satisfies("~> 3.0.1") // true, nil
ver.Satisfies("~> 3.0") // true, nil
ver.Satisfies("~> 2.9") // false, nil
ver.Satisfies("^2.9") // false, nil
// ^ and ~> are the same operator
```

```go
vers := []gosemver.Version{
  gosemver.Version{"", 1, 2, 3, "", ""},
  gosemver.Version{"", 0, 1, 5, "", ""},
  gosemver.Version{"", 1, 0, 0, "", ""},
}

gosemver.FindAll(vers, ">= 1.0.0")
```

## TODO

- more constraints (like [node semver](https://www.npmjs.org/doc/misc/semver.html))
- executable (like, `ls | gosemver sort`, `gosemver inc patch 1.1.0`)
