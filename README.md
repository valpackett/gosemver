# gosemver [![Build Status](https://img.shields.io/travis/myfreeweb/gosemver.svg?style=flat)](https://travis-ci.org/myfreeweb/gosemver) [![Coverage Status](https://img.shields.io/coveralls/myfreeweb/gosemver.svg?style=flat)](https://coveralls.io/r/myfreeweb/gosemver) [![Apache License 2.0](https://img.shields.io/badge/license-Apache%202.0-brightgreen.svg?style=flat)](https://www.tldrlegal.com/l/apache2)

A [Semantic Versioning](http://semver.org) library for the Go programming language.

## Usage

```go
import "github.com/myfreeweb/gosemver"
```

Parsing:

```go
gosemver.ParseVersion("v0.1.0-alpha+build001") // gosemver.Version{"v", 0, 1, 0, "alpha", "build001"}, nil
gosemver.ParseVersion("AAAAAAA") // gosemver.Version{"", 0, 0, 0, "", ""}, error

gosemver.ParseVersions([]string{"0.1.0", "0.2.0"}) // []gosemver.Version{{"", 0, 1, 0, "", ""}, {"", 0, 2, 0, "", ""},}, nil
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

gosemver.FindAll(vers, ">= 1.0.0") // []gosemver.Version{gosemver.Version{"", 1, 2, 3, "", ""}, gosemver.Version{"", 1, 0, 0, "", ""},}
gosemver.FindMax(vers, ">= 1.0.0") // gosemver.Version{"", 1, 2, 3, "", ""}, nil
gosemver.FindMax(vers, ">= 999.0.0") // gosemver.Version{"", 0, 0, 0, "", ""}, error
```

## TODO

- more constraints (like [node semver](https://www.npmjs.org/doc/misc/semver.html))
- executable (like, `ls | gosemver sort`, `gosemver inc patch 1.1.0`)

## License

Copyright 2014 Greg V <floatboth@me.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
