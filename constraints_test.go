package gosemver

import (
	"testing"
)

type constrTestCase struct {
	Input  Version
	Op     string
	Constr Constraint
	Result bool
}

var constrTests = []constrTestCase{
	{Version{"", 1, 1, 0, "", ""}, "*", Constraint{}, true},
	{Version{"", 4, 0, 0, "", ""}, "x", Constraint{}, true},

	{Version{"", 1, 1, 0, "", ""}, "", Constraint{1, 1, 0, true}, true},
	{Version{"", 1, 1, 0, "", ""}, "==", Constraint{1, 1, 0, true}, true},
	{Version{"", 1, 1, 0, "", ""}, "", Constraint{0, 1, 1, true}, false},
	{Version{"", 1, 1, 0, "", ""}, "==", Constraint{0, 1, 1, true}, false},

	{Version{"", 1, 1, 0, "", ""}, "^", Constraint{1, 1, 0, false}, true},
	{Version{"", 1, 1, 0, "", ""}, "~>", Constraint{1, 1, 0, false}, true},
	{Version{"", 1, 1, 256, "", ""}, "~>", Constraint{1, 1, 0, false}, true},
	{Version{"", 1, 1, 256, "", ""}, "~>", Constraint{1, 0, 0, false}, true},
	{Version{"", 2, 0, 0, "", ""}, "~>", Constraint{1, 1, 0, false}, false},
	{Version{"", 1, 1, 0, "", ""}, "~>", Constraint{1, 2, 0, false}, false},

	{Version{"", 3, 0, 1, "", ""}, "~>", Constraint{3, 0, 3, true}, false},
	{Version{"", 3, 0, 3, "", ""}, "~>", Constraint{3, 0, 3, true}, true},
	{Version{"", 3, 0, 10, "", ""}, "~>", Constraint{3, 0, 3, true}, true},
	{Version{"", 3, 1, 0, "", ""}, "~>", Constraint{3, 0, 3, true}, false},

	{Version{"", 1, 1, 0, "", ""}, ">", Constraint{1, 0, 99, true}, true},
	{Version{"", 1, 1, 0, "", ""}, ">", Constraint{1, 1, 0, true}, false},
	{Version{"", 1, 1, 0, "", ""}, ">", Constraint{1, 1, 0, false}, false},
	{Version{"", 1, 1, 0, "", ""}, ">", Constraint{1, 0, 0, false}, true},
	{Version{"", 1, 1, 0, "", ""}, ">=", Constraint{1, 0, 0, false}, true},
	{Version{"", 1, 1, 0, "", ""}, ">=", Constraint{1, 1, 0, true}, true},
	{Version{"", 1, 1, 0, "", ""}, "<=", Constraint{1, 1, 0, true}, true},
	{Version{"", 1, 1, 0, "", ""}, "<=", Constraint{2, 0, 0, false}, true},
	{Version{"", 1, 1, 0, "", ""}, "<", Constraint{1, 1, 5, true}, true},
	{Version{"", 1, 1, 0, "", ""}, "<", Constraint{1, 1, 0, true}, false},
	{Version{"", 1, 1, 0, "", ""}, "<", Constraint{1, 2, 0, false}, true},
}

func TestConstraints(t *testing.T) {
	for _, tcase := range constrTests {
		result := tcase.Input.SatisfiesOp(tcase.Op, &tcase.Constr)
		if result != tcase.Result {
			t.Error(
				"For", tcase.Input, tcase.Op, tcase.Constr,
				"expected", tcase.Result,
				"got", result,
			)
		}
	}
}
