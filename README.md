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
ver := Version{"v", 2, 3, 1, "alpha", "build.001"},
ver.String() // "v2.3.1-alpha+build.001"
```

## TODO

- parsing
- executable (like, `ls | gosemver sort`, `gosemver inc patch 1.1.0`)
