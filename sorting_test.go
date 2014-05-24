package gosemver

import (
	"sort"
	"testing"
)

type testCase struct {
	Before []Version
	After  []Version
}

var tests = []testCase{
	{ // Basic ordering
		[]Version{ // Before
			Version{"v", 0, 1, 0, "", ""},
			Version{"v", 1, 1, 1, "", ""},
			Version{"v", 0, 1, 1, "", ""},
			Version{"v", 1, 1, 0, "", ""},
			Version{"v", 0, 2, 0, "", ""},
			Version{"v", 1, 0, 0, "", ""},
			Version{"v", 0, 2, 1, "", ""},
		},
		[]Version{ // After
			Version{"v", 0, 1, 0, "", ""},
			Version{"v", 0, 1, 1, "", ""},
			Version{"v", 0, 2, 0, "", ""},
			Version{"v", 0, 2, 1, "", ""},
			Version{"v", 1, 0, 0, "", ""},
			Version{"v", 1, 1, 0, "", ""},
			Version{"v", 1, 1, 1, "", ""},
		},
	},
	{ // Identifiers
		[]Version{ // Before
			Version{"", 1, 0, 0, "beta", ""},
			Version{"", 1, 0, 0, "", ""},
			Version{"", 1, 0, 0, "alpha", ""},
		},
		[]Version{ // After
			Version{"", 1, 0, 0, "", ""},
			Version{"", 1, 0, 0, "alpha", ""},
			Version{"", 1, 0, 0, "beta", ""},
		},
	},
}

func TestSorting(t *testing.T) {
	for _, tcase := range tests {
		result := make([]Version, len(tcase.Before))
		copy(result, tcase.Before)
		sort.Sort(Versions(result))
		for i, v := range result {
			if v != tcase.After[i] {
				t.Error(
					"For", tcase.Before,
					"expected", tcase.After,
					"got", result,
				)
			}
		}
	}
}
