package gosemver

import (
	"testing"
)

type constrTestCase struct {
	Input  Version
	Constr string
	Result bool
}

var constrTests = []constrTestCase{
	{Version{"", 1, 1, 0, "", ""}, "*", true},
	{Version{"", 0, 0, 0, "", ""}, "x", true},
	{Version{"", 0, 0, 0, "", ""}, "", true},

	{Version{"", 1, 1, 0, "", ""}, "1.1.0", true},
	{Version{"", 1, 1, 0, "", ""}, "== 1.1.0", true},
	{Version{"", 1, 1, 0, "", ""}, "0.1.1", false},
	{Version{"", 1, 1, 0, "", ""}, "== 0.1.1", false},

	{Version{"", 1, 1, 0, "", ""}, "^1.1", true},
	{Version{"", 1, 1, 0, "", ""}, "~> 1.1", true},
	{Version{"", 1, 1, 256, "", ""}, "~> 1.1", true},
	{Version{"", 1, 1, 256, "", ""}, "~> 1.0", true},
	{Version{"", 2, 0, 0, "", ""}, "~> 1.1", false},
	{Version{"", 1, 1, 0, "", ""}, "~> 1.2", false},

	{Version{"", 3, 0, 1, "", ""}, "~> 3.0.3", false},
	{Version{"", 3, 0, 3, "", ""}, "~> 3.0.3", true},
	{Version{"", 3, 0, 10, "", ""}, "~> 3.0.3", true},
	{Version{"", 3, 1, 0, "", ""}, "~> 3.0.3", false},

	{Version{"", 1, 1, 0, "", ""}, "> 1.0.99", true},
	{Version{"", 1, 1, 0, "", ""}, "> 1.1.0", false},
	{Version{"", 1, 1, 0, "", ""}, "> 1.1", false},
	{Version{"", 1, 1, 0, "", ""}, "> 1.0", true},
	{Version{"", 1, 1, 0, "", ""}, ">= 1.0", true},
	{Version{"", 1, 1, 0, "", ""}, ">= 1.1.0", true},
	{Version{"", 1, 1, 0, "", ""}, "<= 1.1.0", true},
	{Version{"", 1, 1, 0, "", ""}, "<= 2.0", true},
	{Version{"", 1, 1, 0, "", ""}, "< 1.1.5", true},
	{Version{"", 1, 1, 0, "", ""}, "< 1.1.0", false},
	{Version{"", 1, 1, 0, "", ""}, "< 1.2", true},
}

func TestConstraints(t *testing.T) {
	for _, tcase := range constrTests {
		result := tcase.Input.Satisfies(tcase.Constr)
		if result != tcase.Result {
			t.Error(
				"For", tcase.Input, tcase.Constr,
				"expected", tcase.Result,
				"got", result,
			)
		}
	}
}
