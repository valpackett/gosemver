# gosemver

A [Semantic Versioning](http://semver.org) library for the Go programming language.

## Usage

```go
import "github.com/myfreeweb/gosemver"
```

Parsing:

```go
gosemver.parseVersion("v0.1.0-alpha+build001") // Version{"v", 0, 1, 0, "alpha", "build001"}
```

Sorting:

```go
import "sort"

vers := []gosemver.Version{
  {"v", 1, 0, 0, "", ""},
  {"v", 0, 1, 0, "", ""},
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
ver.Satisfies("*") // true
ver.Satisfies("== 3.0.3") // true
ver.Satisfies(">= 3.0.1") // true
ver.Satisfies(">= 3.0") // true
ver.Satisfies("> 3.0.0") // true
ver.Satisfies("~> 3.0.4") // false
ver.Satisfies("~> 3.0.1") // true
ver.Satisfies("~> 3.0") // true
ver.Satisfies("~> 2.9") // false
ver.Satisfies("^2.9") // false
// ^ and ~> are the same operator
```

## TODO

- more constraints (like [node semver](https://www.npmjs.org/doc/misc/semver.html))
- executable (like, `ls | gosemver sort`, `gosemver inc patch 1.1.0`)
