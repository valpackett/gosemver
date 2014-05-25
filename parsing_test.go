package gosemver

import (
	"testing"
)

type parseConstrTestCase struct {
	Input        string
	ResultOp     string
	ResultConstr Constraint
}

var parseConstrTests = []parseConstrTestCase{
	{"0.1.0", "", Constraint{0, 1, 0, true}},
	{"0.1", "", Constraint{0, 1, 0, false}},
	{"~> 128.256", "~>", Constraint{128, 256, 0, false}},
	{"~>128.256", "~>", Constraint{128, 256, 0, false}},
	{"^128.256.512", "^", Constraint{128, 256, 512, true}},
}

func TestParseConstraints(t *testing.T) {
	for _, tcase := range parseConstrTests {
		resultOp, resultConstr := parseConstraint(tcase.Input)
		if resultOp != tcase.ResultOp || *resultConstr != tcase.ResultConstr {
			t.Error(
				"For", tcase.Input,
				"expected", tcase.ResultOp, tcase.ResultConstr,
				"got", resultOp, *resultConstr,
			)
		}
	}
}
