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
		resultOp, resultConstr, _ := parseConstraint(tcase.Input)
		if resultOp != tcase.ResultOp || resultConstr != tcase.ResultConstr {
			t.Error(
				"For", tcase.Input,
				"expected", tcase.ResultOp, tcase.ResultConstr,
				"got", resultOp, resultConstr,
			)
		}
	}
}

type parseVersionTestCase struct {
	Input  string
	Result Version
}

var parseVersionTests = []parseVersionTestCase{
	{"0.1.0", Version{"", 0, 1, 0, "", ""}},
	{"v0.1.0", Version{"v", 0, 1, 0, "", ""}},
	{"v0.1.0-alpha.1", Version{"v", 0, 1, 0, "alpha.1", ""}},
	{"=0.1.0-beta+build.0001", Version{"=", 0, 1, 0, "beta", "build.0001"}},
	{"версия1024.2048.4096-βββ+¾", Version{"версия", 1024, 2048, 4096, "βββ", "¾"}}, // semver.org actually restricts identifiers/buildmeta to [a-zA-Z0-9-], who cares though
}

func TestParseVersions(t *testing.T) {
	for _, tcase := range parseVersionTests {
		result, _ := parseVersion(tcase.Input)
		if result != tcase.Result {
			t.Error(
				"For", tcase.Input,
				"expected", tcase.Result,
				"got", result,
			)
		}
	}
}
