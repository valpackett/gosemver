# gosemver

A [Semantic Versioning](http://semver.org) library for the Go programming language.

## Usage

```go
import "github.com/myfreeweb/gosemver"
```

Sorting:

```go
import "sort"

vers := []Version{
  Version{"v", 1, 0, 0, "", ""},
  Version{"v", 0, 1, 0, "", ""},
}

sort.Sort(Versions(vers))
```

Output:

```go
ver := Version{"v", 2, 3, 1, "alpha", "build.001"}
ver.String() // "v2.3.1-alpha+build.001"
```

Constraints:

```go
ver := Version{"", 3, 0, 3, "", ""}
ver.Satisfies("*") // true
ver.Satisfies("~> 3.0.1") // true
ver.Satisfies("~> 3.0") // true
ver.Satisfies("~> 2.9") // false
ver.Satisfies("^2.9") // false
// ^ and ~> are the same operator
```

## TODO

- parsing
- more constaints (like [node semver](https://www.npmjs.org/doc/misc/semver.html))
- executable (like, `ls | gosemver sort`, `gosemver inc patch 1.1.0`)
